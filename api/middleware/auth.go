package middleware

import (
	"github.com/Waifu-pics/waifu-api/api"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/labstack/echo"
)

// AuthPayload : JWT auth payload
type AuthPayload struct {
	jwt.Payload
	Identifier string `json:"Identifier,omitempty"`
}

// Auth : middleware instance for authenticatio
type Auth struct {
	Jwt string
}

// VerifySubtle : JWT authentication without errors
func (m Auth) VerifySubtle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("authbool", false)

		token, err := c.Cookie("auth-token")
		if err != nil {
			return next(c)
		}

		var pl AuthPayload
		_, err = jwt.Verify([]byte(token.Value), jwt.NewHS256([]byte(m.Jwt)), &pl)
		if err != nil {
			return next(c)
		}

		c.Set("authbool", true)

		return next(c)
	}
}

// Verify : JWT authentication
func (m Auth) Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := c.Cookie("auth-token")
		if err != nil {
			c.JSON(401, api.Basic{Message: "No auth token"})
			return nil
		}

		var pl AuthPayload
		_, err = jwt.Verify([]byte(token.Value), jwt.NewHS256([]byte(m.Jwt)), &pl)
		if err != nil {
			c.JSON(400, api.Basic{Message: "Invalid auth token"})
			return nil
		}

		return next(c)
	}
}
