package fun

import (
	"time"

	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/api/middleware"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo"
)

type route api.Options

// NewRouter : create new router for image endpoint
func NewRouter(options api.Options, c *echo.Group) {
	route := route(options)
	// Ratelimit
	c.POST("/gen", route.GenerateMeme, middleware.LimitMiddleware(tollbooth.NewLimiter(0.5, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Minute,
	})))
}
