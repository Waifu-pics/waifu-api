package image

import (
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/labstack/echo"
)

// Route : image endpoints
type Route struct {
	Options  api.Options
	Nsfw     bool
	Endpoint string
}

// NewRouter : create new router for image endpoint
func NewRouter(options api.Options, c *echo.Group) {
	for _, endpoint := range options.Config.Endpoints.Sfw {
		route := Route{
			Options:  options,
			Nsfw:     false,
			Endpoint: endpoint,
		}

		c.GET("/sfw/"+endpoint, route.GetImage)
		c.POST("/many/sfw/"+endpoint, route.GetManyImage)
	}

	for _, endpoint := range options.Config.Endpoints.Nsfw {
		route := Route{
			Options:  options,
			Nsfw:     true,
			Endpoint: endpoint,
		}

		c.GET("/nsfw/"+endpoint, route.GetImage)
		c.POST("/many/nsfw/"+endpoint, route.GetManyImage)
	}
}
