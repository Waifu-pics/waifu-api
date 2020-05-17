package app

import (
	"net/http"

	"github.com/gorilla/mux"

	api "waifu.pics/app/API"
	"waifu.pics/util"
)

// Router : Init router function
func Router(mux *mux.Router, config util.Config) *mux.Router {

	// Execute this loop for every endpoint in config
	for _, endP := range config.ENDPOINTS {
		endpoint := endP // Evaluates instantly
		api.SingleImagePoint(mux, endpoint, config)
		api.ManyImagePoint(mux, endpoint, config)
		Grid(mux, endpoint, config)
	}

	// Rest of front end
	Docs(mux, config)
	api.UploadHandle(mux, config)

	// Other important things
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public/static/"))))
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(Error404).GetHandler()

	return mux
}
