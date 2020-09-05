package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Riku32/waifu.pics/app"
	"github.com/Riku32/waifu.pics/util/admin"
	"github.com/Riku32/waifu.pics/util/config"
	"github.com/Riku32/waifu.pics/util/database"
	"github.com/Riku32/waifu.pics/util/file"
	"github.com/Riku32/waifu.pics/util/static"
)

func main() {
	dev := flag.Bool("dev", false, "developer mode")
	newuser := flag.Bool("newuser", false, "create an administrator")
	flag.Parse()
	static.Dev = *dev

	var cfg config.Config // Config paths different
	if !static.Dev {
		cfg = config.LoadConfig("./config.yml")
	} else {
		cfg = config.LoadConfig("../../config.yml")
	}

	db := database.InitSQL(cfg)

	// Admin creation argument
	if *newuser {
		admin.CreateAdmin(db)
		return
	}

	file.InitS3(cfg)

	err := http.ListenAndServe(":"+cfg.Port, app.Router(cfg, db))

	if err != nil {
		log.Fatal("Unable to start the web server!")
	}
}
