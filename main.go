package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"waifu.pics/app"
	"waifu.pics/util"
)

func main() {
	config := util.LoadConfig("config.json")
	database := util.DbDriver(config)
	http.ListenAndServe(":"+config.PORT, app.Router(mux.NewRouter(), &database, config))
}
