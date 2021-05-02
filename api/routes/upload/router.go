package upload

import (
	"github.com/Waifu-pics/waifu-api/api"
	"github.com/Waifu-pics/waifu-api/api/middleware"
	"github.com/labstack/echo"
)

// Route : route object
type Route api.Options

// NewRouter : create new router for upload endpoint
func NewRouter(options api.Options, c *echo.Group) {
	route := Route(options)
	auth := middleware.Auth{
		Jwt: options.Config.Web.Jwt,
	}
	c.Group("", auth.VerifySubtle).POST("/upload", route.UploadHandle)
}
