package app

import (
	"net/http"
	"waifu.pics/util/config"

	"go.mongodb.org/mongo-driver/mongo"
	"waifu.pics/app/Views"

	"github.com/gorilla/mux"

	api "waifu.pics/app/API"
)

// Router : Init router function
func Router(mux *mux.Router, config config.Config, database *mongo.Database) *mux.Router {
	front := Views.Front{Endpoints: config.ENDPOINTS}
	endpoints := &api.API{Config: config, Database: database}
	admin := &Views.Admin{Database: database}

	// Execute this loop for every endpoint in config
	for _, endP := range config.ENDPOINTS {
		endpoint := endP // Evaluates instantly

		apiMulti := api.Multi{Endpoint: endpoint, API: endpoints}     // Views Multi function
		viewsMulti := Views.Multi{Endpoint: endpoint, Config: config} // API Multi function

		mux.HandleFunc("/api/"+endpoint, apiMulti.GetImage).Methods("GET")
		mux.HandleFunc("/api/many/"+endpoint, apiMulti.GetManyImage).Methods("POST")

		switch endpoint {
		case "sfw":
			mux.HandleFunc("/", viewsMulti.Grid)
		default:
			mux.HandleFunc("/"+endpoint, viewsMulti.Grid)
		}
	}

	// Front end
	mux.HandleFunc("/docs", front.Docs)
	mux.HandleFunc("/upload", front.UploadFront)
	mux.HandleFunc("/admin", admin.AdminPage)
	mux.HandleFunc("/pages", front.Pages)

	// Api stuff
	mux.HandleFunc("/api/upload", endpoints.UploadHandle).Methods("POST")
	mux.HandleFunc("/api/admin/login", endpoints.AdminLogin).Methods("POST")
	mux.HandleFunc("/api/admin/token", endpoints.AdminVerify).Methods("POST")
	mux.HandleFunc("/api/admin/list", endpoints.ListFile).Methods("POST")
	mux.HandleFunc("/api/admin/verify", endpoints.VerifyFile).Methods("POST")
	mux.HandleFunc("/api/endpoints", endpoints.GetEndpoints).Methods("GET")

	// Other important things
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public/static/"))))
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(Views.Error404).GetHandler()

	return mux
}
