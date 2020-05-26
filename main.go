package main

import (
	"log"
	"net/http"
	"os"
	"waifu.pics/util/admin"
	"waifu.pics/util/config"
	"waifu.pics/util/file"

	"github.com/gorilla/mux"
	"waifu.pics/app"
	"waifu.pics/util/database"
)

func main() {
	cfg := config.LoadConfig("config.json")
	db := database.InitDB(cfg)

	// Admin creation argument
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "createAdmin" {
			admin.CreateAdmin(db)
			return
		}
	}

	file.InitS3(cfg)

	err := http.ListenAndServe(":"+cfg.PORT, app.Router(mux.NewRouter(), cfg, db))

	if err != nil {
		log.Fatal("Unable to start the web server!")
	}
}
