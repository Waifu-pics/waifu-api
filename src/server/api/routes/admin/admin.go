package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/api/middleware"
	"github.com/alexedwards/argon2id"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/labstack/echo"
)

// Login : log in to the API
func (i Route) Login(c echo.Context) error {
	if c.Get("authbool").(bool) {
		return c.JSON(202, api.Basic{Message: "Already logged in"})
	}

	body := new(Credentials)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	hash, err := i.Database.GetAdminHash(body.Username)
	if err != nil {
		return c.JSON(400, api.Basic{Message: "Incorrect credentials!"})
	}

	valid, err := argon2id.ComparePasswordAndHash(body.Password, hash)
	if err != nil {
		return c.JSON(500, api.Basic{Message: api.ErrServer})
	}

	if !valid {
		return c.JSON(400, api.Basic{Message: "Incorrect credentials!"})
	}

	payload := middleware.AuthPayload{
		Payload: jwt.Payload{
			Issuer:         "waifu.pics",
			ExpirationTime: jwt.NumericDate(time.Now().Add(24 * 30 * 12 * time.Hour)),
			IssuedAt:       jwt.NumericDate(time.Now()),
		},
		Identifier: body.Username,
	}

	jwtoken, err := jwt.Sign(payload, jwt.NewHS256([]byte(i.Config.Web.Jwt)))
	if err != nil {
		return c.JSON(500, api.Basic{Message: api.ErrServer})
	}

	cookie := http.Cookie{
		Name:    "auth-token",
		Value:   string(jwtoken),
		Expires: time.Now().Add(60 * time.Minute),
		Path:    "/",
	}

	c.SetCookie(&cookie)
	return c.JSON(201, api.Basic{Message: "You have been logged in!"})
}

// ListFile : listing files
func (i Route) ListFile(c echo.Context) error {
	body := new(Search)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	if !api.CheckValid(body.Endpoint, body.Nsfw, i.Config) {
		return c.JSON(400, api.Basic{Message: "invalid type"})
	}

	files, err := i.Database.GetFilesAdmin(body.Endpoint, body.Query, body.Verified, body.Nsfw)
	if err != nil {
		return c.JSON(500, api.Basic{Message: api.ErrServer})
	}

	var response struct {
		Files []File `json:"files"`
	}

	for _, v := range files {
		response.Files = append(response.Files, File{
			Name: v,
			URL:  i.Config.Web.Cdn + v,
		})
	}

	return c.JSON(200, response)
}

// VerifyFile : Verifying user uploads
func (i Route) VerifyFile(c echo.Context) error {
	body := new(Filelist)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	var errcount int

	for _, v := range body.Files {
		err := i.Database.VerifyFile(v)
		if err != nil {
			errcount++
		}
	}

	return c.JSON(200, api.Basic{Message: fmt.Sprintf("Files have been verified with %d errors", errcount)})
}

// DeleteFile : delete a file from the API
func (i Route) DeleteFile(c echo.Context) error {
	body := new(Filelist)
	if err := c.Bind(body); err != nil {
		return c.JSON(400, api.Basic{Message: api.ErrInvalidJSON})
	}

	var errcount int

	for _, v := range body.Files {
		err := i.Database.DeleteFile(v)
		if err != nil {
			errcount++
		}
		if err := i.S3.DeleteFile(v); err != nil {
			errcount++
		}
	}

	return c.JSON(200, api.Basic{Message: fmt.Sprintf("Files have been deleted with %d errors", errcount)})
}
