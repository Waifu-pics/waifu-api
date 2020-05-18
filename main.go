package main

import (
	"net/http"

	"waifu.pics/util/database"
	"waifu.pics/util/file"

	"github.com/gorilla/mux"
	"waifu.pics/app"
	"waifu.pics/util"
)

func main() {
	config := util.LoadConfig("config.json")
	db := database.InitDB(config)
	file.InitS3(config)
	http.ListenAndServe(":"+config.PORT, app.Router(mux.NewRouter(), config, db))
}
