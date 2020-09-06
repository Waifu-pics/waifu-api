package admin

import (
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/middleware"
	"github.com/labstack/echo"
)

// Route : route object
type Route api.Options

// NewRouter : create new router for admin endpoint
func NewRouter(options api.Options, c *echo.Group) {
	routes := Route(options)

	auth := middleware.Auth{
		Jwt: options.Config.Web.Jwt,
	}

	admin := c.Group("/admin")
	admin.Group("", auth.VerifySubtle).POST("/login", routes.Login)

	authreq := admin.Group("", auth.Verify)
	authreq.POST("/verify", routes.VerifyFile)
	authreq.POST("/delete", routes.DeleteFile)
	authreq.POST("/list", routes.ListFile)
}
