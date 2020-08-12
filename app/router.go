package app

import (
	api "waifu.pics/app/API"
	"waifu.pics/app/auth"
	"waifu.pics/util/config"
	"waifu.pics/util/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router : Init router function
func Router(config config.Config, database database.Database) *chi.Mux {
	r := chi.NewRouter()
	endpoints := &api.API{Config: config, Database: database}
	mw := &auth.Middleware{Config: config}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}))

	r.Get("/api/endpoints", endpoints.GetEndpoints)

	r.Group(func(r chi.Router) {
		r.Use(mw.VerifySubtle)

		r.Post("/api/login", endpoints.Login)
		r.Post("/api/upload", endpoints.UploadHandle)
	})

	r.Group(func(r chi.Router) {
		r.Use(mw.Verify)

		r.Post("/api/admin/verify", endpoints.VerifyFile)
		r.Post("/api/admin/delete", endpoints.DeleteFile)
		r.Post("/api/admin/list", endpoints.ListFile)
	})

	// Execute this loop for every endpoint in config
	for _, endP := range config.ENDPOINTS {
		endpoint := endP // Evaluates instantly

		apiMulti := api.Multi{Endpoint: endpoint, API: endpoints}

		r.Get("/api/"+endpoint, apiMulti.GetImage)
		r.Post("/api/many/"+endpoint, apiMulti.GetManyImage)
	}

	return r
}
