package middleware

import (
	"github.com/Waifu-pics/waifu-api/api"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo"
)

// LimitMiddleware : ratelimit layer for tollbooth
func LimitMiddleware(lmt *limiter.Limiter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			httpError := tollbooth.LimitByRequest(lmt, c.Response(), c.Request())
			if httpError != nil {
				return c.JSON(httpError.StatusCode, api.Basic{Message: "You are being rate limited"})
			}
			return next(c)
		})
	}
}
