package router

import (
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/api/routes/admin"
	"github.com/Riku32/waifu.pics/src/api/routes/image"
	"github.com/Riku32/waifu.pics/src/api/routes/info"
	"github.com/Riku32/waifu.pics/src/api/routes/upload"
	"github.com/Riku32/waifu.pics/src/static"
	"github.com/labstack/echo"
)

// New : initialize router
func New(options api.Options) {
	e := echo.New()

	// Serve frontend, disabled in dev mode
	if !static.Dev {
		e.Static("/", "dist/")
		echo.NotFoundHandler = func(c echo.Context) error {
			return c.File("./dist/index.html")
		}
	}

	api := e.Group("/api")
	image.NewRouter(options, api)
	admin.NewRouter(options, api)
	upload.NewRouter(options, api)
	info.NewRouter(options, api)
	//fun.NewRouter(options, api)

	api.GET("/endpoints", func(c echo.Context) error {
		return c.JSON(200, options.Config.Endpoints)
	})

	e.Logger.Fatal(e.Start(":" + options.Config.Port))
}
