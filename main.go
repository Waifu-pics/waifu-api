package main

import (
	"log"
	"net/http"
	"os"

	"waifu.pics/util/admin"
	"waifu.pics/util/config"
	"waifu.pics/util/file"

	"waifu.pics/app"
	"waifu.pics/util/database"
)

func main() {
	cfg := config.LoadConfig("config.json")
	db := database.InitSQL(cfg)

	// Admin creation argument
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "createadmin" {
			admin.CreateAdmin(db)
			return
		}
	}

	file.InitS3(cfg)

	err := http.ListenAndServe(":"+cfg.PORT, app.Router(cfg, db))

	if err != nil {
		log.Fatal("Unable to start the web server!")
	}
}
