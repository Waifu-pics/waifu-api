package router

import (
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/api/routes/admin"
	"github.com/Riku32/waifu.pics/src/api/routes/image"
	"github.com/Riku32/waifu.pics/src/api/routes/info"
	"github.com/Riku32/waifu.pics/src/api/routes/upload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New : initialize router
func New(options api.Options) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	api := e.Group("/api")
	image.NewRouter(options, api)
	admin.NewRouter(options, api)
	upload.NewRouter(options, api)
	info.NewRouter(options, api)

	api.GET("/endpoints", func(c echo.Context) error {
		return c.JSON(200, options.Config.Endpoints)
	})

	e.Logger.Fatal(e.Start(":" + options.Config.Port))
}
