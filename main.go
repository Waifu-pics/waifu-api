package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"waifu.pics/app"
	"waifu.pics/util"
)

func main() {
	config := util.LoadConfig("config.json")
	database := util.DbDriver(string(config.DB))
	http.ListenAndServe(":"+config.PORT, app.Router(mux.NewRouter(), database, config))
}
