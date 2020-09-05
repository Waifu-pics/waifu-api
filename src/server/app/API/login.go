package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Riku32/waifu.pics/src/server/app/auth"
	"github.com/Riku32/waifu.pics/src/server/util/crypto"
	"github.com/Riku32/waifu.pics/src/server/util/web"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gorilla/context"
)

func (api API) Login(w http.ResponseWriter, r *http.Request) {
	if context.Get(r, "authbool").(bool) {
		web.SendJSON(w, 202, web.NewMessage("already logged in"))
		return
	}

	var res struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage(invalidjson))
		return
	}

	hash, err := api.Database.GetAdminHash(res.Username)
	if err != nil {
		web.SendJSON(w, 400, web.NewMessage("Incorrect credentials!"))
		return
	}

	valid, err := crypto.ComparePassword(res.Password, hash)
	if err != nil {
		web.SendJSON(w, 500, web.NewMessage(servererror))
		return
	}

	if !valid {
		web.SendJSON(w, 400, web.NewMessage("Incorrect credentials!"))
		return
	}

	payload := auth.Payload{
		Payload: jwt.Payload{
			Issuer:         "waifu.pics",
			ExpirationTime: jwt.NumericDate(time.Now().Add(24 * 30 * 12 * time.Hour)),
			IssuedAt:       jwt.NumericDate(time.Now()),
		},
		Identifier: res.Username,
	}

	jwtoken, err := jwt.Sign(payload, jwt.NewHS256([]byte(api.Config.Web.Jwt)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:    "auth-token",
		Value:   string(jwtoken),
		Expires: time.Now().Add(60 * time.Minute),
		Path:    "/",
	}

	http.SetCookie(w, &cookie)
	web.SendJSON(w, 201, web.NewMessage("You have been logged in!"))
	return
}
