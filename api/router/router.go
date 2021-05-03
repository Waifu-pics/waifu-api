package router

import (
	"net/http"

	"github.com/Waifu-pics/waifu-api/api"
	"github.com/Waifu-pics/waifu-api/api/routes/admin"
	"github.com/Waifu-pics/waifu-api/api/routes/image"
	"github.com/Waifu-pics/waifu-api/api/routes/info"
	"github.com/Waifu-pics/waifu-api/api/routes/upload"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New : initialize router
func New(options api.Options) {
	e := echo.New()

	api := e.Group("") // Root URL for the API location
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"https://waifu.pics"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	image.NewRouter(options, api)
	admin.NewRouter(options, api)
	upload.NewRouter(options, api)
	info.NewRouter(options, api)

	api.GET("/endpoints", func(c echo.Context) error {
		return c.JSON(200, options.Config.Endpoints)
	})

	e.Logger.Fatal(e.Start(":" + options.Config.Port))
}
