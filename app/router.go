package app

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"waifu.pics/app/Views"

	"github.com/gorilla/mux"

	api "waifu.pics/app/API"
	"waifu.pics/util"
)

// Router : Init router function
func Router(mux *mux.Router, config util.Config, database *mongo.Database) *mux.Router {
	front := Views.Front{config.ENDPOINTS}
	endpoints := &api.API{Config: config, Database: database}

	// Execute this loop for every endpoint in config
	for _, endP := range config.ENDPOINTS {
		endpoint := endP // Evaluates instantly

		apiMulti := api.Multi{Endpoint: endpoint, API: endpoints}     // Views Multi function
		viewsMulti := Views.Multi{Endpoint: endpoint, Config: config} // API Multi function

		mux.HandleFunc("/api/"+endpoint, apiMulti.GetImage).Methods("GET")
		mux.HandleFunc("/api/many/"+endpoint, apiMulti.GetManyImage).Methods("POST")

		// If endpoint is sfw then use /
		if endpoint == "sfw" {
			mux.HandleFunc("/", viewsMulti.Grid)
		} else {
			mux.HandleFunc("/"+endpoint, viewsMulti.Grid)
		}
	}

	// Front end
	mux.HandleFunc("/docs", front.Docs)
	mux.HandleFunc("/upload", front.UploadFront)
	mux.HandleFunc("/admin", Views.AdminLogin)
	mux.HandleFunc("/admin/dash", Views.AdminDash)

	// Api stuff
	mux.HandleFunc("/api/upload", endpoints.UploadHandle).Methods("POST")
	mux.HandleFunc("/api/admin/login", endpoints.AdminLogin).Methods("POST")
	mux.HandleFunc("/api/admin/verifytoken", endpoints.AdminVerify).Methods("POST")
	mux.HandleFunc("/api/admin/list", endpoints.ListFile).Methods("POST")
	mux.HandleFunc("/api/admin/verify", endpoints.VerifyFile).Methods("POST")

	// Other important things
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public/static/"))))
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(Views.Error404).GetHandler()

	return mux
}
