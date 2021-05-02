package info

import (
	"github.com/Waifu-pics/waifu-api/api"
	"github.com/labstack/echo"
)

// Route : route object
type route api.Options

// NewRouter : create new router for upload endpoint
func NewRouter(options api.Options, c *echo.Group) {
	route := route(options)
	c.GET("/recent", route.RecentFiles)
}
