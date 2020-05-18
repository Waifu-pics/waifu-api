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

	// Front end
	Docs(mux, config)
	UploadFront(mux, config)
	mux.HandleFunc("/admin", AdminLogin)
	mux.HandleFunc("/admin/dash", AdminDash)

	// Api stuff
	api.UploadHandle(mux, config)
	mux.HandleFunc("/api/admin/login", api.AdminLogin).Methods("POST")
	mux.HandleFunc("/api/admin/verifytoken", api.AdminVerify).Methods("POST")
	mux.HandleFunc("/api/admin/list", api.ListFile).Methods("POST")
	api.VerifyFile(mux, config)

	// Other important things
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public/static/"))))
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(Error404).GetHandler()

	return mux
}
