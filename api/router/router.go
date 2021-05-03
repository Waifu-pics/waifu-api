package router

import (
	"github.com/Waifu-pics/waifu-api/api"
	"github.com/Waifu-pics/waifu-api/api/routes/admin"
	"github.com/Waifu-pics/waifu-api/api/routes/image"
	"github.com/Waifu-pics/waifu-api/api/routes/info"
	"github.com/Waifu-pics/waifu-api/api/routes/upload"
	"github.com/labstack/echo"
)

// New : initialize router
func New(options api.Options) {
	e := echo.New()

	api := e.Group("") // Root URL for the API location
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return next(ctx)
		}
	})

	image.NewRouter(options, api)
	admin.NewRouter(options, api)
	upload.NewRouter(options, api)
	info.NewRouter(options, api)

	api.GET("/endpoints", func(c echo.Context) error {
		return c.JSON(200, options.Config.Endpoints)
	})

	e.Logger.Fatal(e.Start(":" + options.Config.Port))
}
