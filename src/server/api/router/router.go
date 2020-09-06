package router

import (
	s3simple "github.com/Riku32/s3-simple"
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/api/routes/admin"
	"github.com/Riku32/waifu.pics/src/api/routes/image"
	"github.com/Riku32/waifu.pics/src/api/routes/upload"
	"github.com/Riku32/waifu.pics/src/config"
	"github.com/Riku32/waifu.pics/src/database"
	"github.com/Riku32/waifu.pics/src/static"
	"github.com/labstack/echo"
)

// New : initialize router
func New(conf config.Config, database database.Database, s3 *s3simple.Session) {
	e := echo.New()

	// Serve frontend, disabled in dev mode
	if !static.Dev {
		e.Static("/", "dist/")
		echo.NotFoundHandler = func(c echo.Context) error {
			return c.File("./dist/index.html")
		}
	}

	options := api.Options{
		Database: database,
		Config:   conf,
		S3:       s3,
	}

	api := e.Group("/api")
	image.NewRouter(options, api)
	admin.NewRouter(options, api)
	upload.NewRouter(options, api)

	api.GET("/endpoints", func(c echo.Context) error {
		return c.JSON(200, conf.Endpoints)
	})

	e.Logger.Fatal(e.Start(":" + conf.Port))
}
