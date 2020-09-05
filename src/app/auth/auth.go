package auth

import (
	"net/http"

	"github.com/Riku32/waifu.pics/src/util/config"
	"github.com/Riku32/waifu.pics/src/util/web"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gorilla/context"
)

// Payload : JWT auth payload
type Payload struct {
	jwt.Payload
	Identifier string `json:"Identifier,omitempty"`
}

// Middleware : middleware instance for auth
type Middleware struct {
	Config config.Config
}

// VerifySubtle : JWT authentication without errors
func (m Middleware) VerifySubtle(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "authbool", false)
		token, err := web.GetCookie(r.Cookies(), "auth-token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		var pl Payload
		_, err = jwt.Verify([]byte(token), jwt.NewHS256([]byte(m.Config.Web.Jwt)), &pl)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		context.Set(r, "authbool", true)
		context.Set(r, "auth", pl)

		next.ServeHTTP(w, r)
		return
	}

	return http.HandlerFunc(fn)
}

// Verify : JWT authentication
func (m Middleware) Verify(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := web.GetCookie(r.Cookies(), "auth-token")
		if err != nil {
			web.WriteResp(w, 401, "No auth token")
			r.Body.Close()
			return
		}

		var pl Payload
		_, err = jwt.Verify([]byte(token), jwt.NewHS256([]byte(m.Config.Web.Jwt)), &pl)
		if err != nil {
			web.WriteResp(w, 400, "Invalid auth token")
			r.Body.Close()
			return
		}

		context.Set(r, "auth", pl)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
