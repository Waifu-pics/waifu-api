package upload

import (
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/labstack/echo"
)

// Route : route object
type Route api.Options

// NewRouter : create new router for upload endpoint
func NewRouter(options api.Options, c *echo.Group) {
	route := Route(options)
	c.POST("/upload", route.UploadHandle)
}
