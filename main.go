package main

import (
	"net/http"
	"waifu.pics/util/config"

	"waifu.pics/util/database"
	"waifu.pics/util/file"

	"github.com/gorilla/mux"
	"waifu.pics/app"
)

func main() {
	cfg := config.LoadConfig("config.json")
	db := database.InitDB(cfg)
	file.InitS3(cfg)
	http.ListenAndServe(":"+cfg.PORT, app.Router(mux.NewRouter(), cfg, db))
}
