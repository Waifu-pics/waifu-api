package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gorilla/context"
	"waifu.pics/app/auth"
	"waifu.pics/util/crypto"
	"waifu.pics/util/web"
)

func (api API) Login(w http.ResponseWriter, r *http.Request) {
	if context.Get(r, "authbool").(bool) {
		web.WriteResp(w, 202, "Already logged in!")
		return
	}

	var res struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		web.WriteResp(w, 400, "Invalid JSON")
		return
	}

	hash, err := api.Database.GetAdminHash(res.Username)
	if err != nil {
		web.WriteResp(w, 400, "Incorrect credentials!")
		return
	}

	valid, err := crypto.ComparePassword(res.Password, hash)
	if err != nil {
		web.WriteResp(w, 400, "Error!")
		return
	}

	if !valid {
		web.WriteResp(w, 400, "Incorrect credentials!")
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

	jwtoken, err := jwt.Sign(payload, jwt.NewHS256([]byte(api.Config.JWT)))
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
	web.WriteResp(w, 201, "You have been logged in!")
	return
}
