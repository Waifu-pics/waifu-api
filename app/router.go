package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"waifu.pics/app/API"
	"waifu.pics/util"
)

// Router : Init router function
func Router(mux *mux.Router, config util.Config) *mux.Router {

	// Execute this loop for every endpoint in config
	for _, endP := range config.ENDPOINTS {
		endpoint := endP // Evaluates instantly
		API.SingleImagePoint(mux, endpoint, config)
		API.ManyImagePoint(mux, endpoint, config)
		Grid(mux, endpoint, config)
	}

	Docs(mux, config)

	// Static file server
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("external/static/"))))

	return mux
}
