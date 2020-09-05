package app

import (
	"fmt"
	"net/http"
	"strings"

	api "github.com/Riku32/waifu.pics/src/app/API"
	"github.com/Riku32/waifu.pics/src/app/auth"
	"github.com/Riku32/waifu.pics/src/util/config"
	"github.com/Riku32/waifu.pics/src/util/database"
	"github.com/Riku32/waifu.pics/src/util/static"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router : Init router function
func Router(config config.Config, database database.Database) *chi.Mux {
	r := chi.NewRouter()
	endpoints := &api.API{Config: config, Database: database}
	mw := &auth.Middleware{Config: config}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
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

	// Image Routes
	for _, endpoint := range config.Endpoints.Sfw {
		apim := api.Multi{
			API:      endpoints,
			Nsfw:     false,
			Endpoint: endpoint,
		}

		r.Get("/api/sfw/"+endpoint, apim.GetImage)
		r.Post("/api/many/sfw/"+endpoint, apim.GetManyImage)
	}

	for _, endpoint := range config.Endpoints.Nsfw {
		apim := api.Multi{
			API:      endpoints,
			Nsfw:     true,
			Endpoint: endpoint,
		}

		r.Get("/api/nsfw/"+endpoint, apim.GetImage)
		r.Post("/api/many/nsfw/"+endpoint, apim.GetManyImage)
	}

	if static.Dev {
		ServeDir(r, "/", "./dist/")

		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./dist/index.html")
		})
	}

	return r
}

// ServeDir : serve directory
func ServeDir(router *chi.Mux, path, directory string) {
	rpath := http.Dir(fmt.Sprintf("%s/*", path))
	filepath := http.Dir(directory)
	router.Get(string(rpath), func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filepath))
		fs.ServeHTTP(w, r)
	})
}
