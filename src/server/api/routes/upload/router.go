package upload

import (
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/middleware"
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
