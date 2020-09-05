package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Riku32/waifu.pics/src/app"
	"github.com/Riku32/waifu.pics/src/util/admin"
	"github.com/Riku32/waifu.pics/src/util/config"
	"github.com/Riku32/waifu.pics/src/util/database"
	"github.com/Riku32/waifu.pics/src/util/file"
	"github.com/Riku32/waifu.pics/src/util/static"
)

func main() {
	serve := flag.Bool("serve", false, "serving gui")
	newuser := flag.Bool("newuser", false, "create an administrator")
	flag.Parse()
	static.Serve = *serve

	cfg := config.LoadConfig("./config.yml")

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

	log.Printf("Started webserver on port %s\n", cfg.Port)
}
