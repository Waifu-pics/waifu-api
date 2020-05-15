package app

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"waifu.pics/util"
)

type grid struct {
	URL      string
	Endpoint string
}

// Grid : This is the grid page initializer for every endpoint
func Grid(mux *mux.Router, endpoint string, conf util.Config) {
	p := grid{URL: conf.URL, Endpoint: endpoint}
	// Setting up all templates
	t := template.Must(template.ParseFiles(
		"public/templates/grid.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	// This is separate because sfw should be on index
	if endpoint == "sfw" {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			t.ExecuteTemplate(w, "grid", p)
			defer r.Body.Close()
		}).Methods("GET")
	}

	mux.HandleFunc("/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "grid", p)
		defer r.Body.Close()
	}).Methods("GET")
}

type docs struct {
	Endpoints []string
}

// Docs : This is the api page
func Docs(mux *mux.Router, conf util.Config) {
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		data := docs{Endpoints: conf.ENDPOINTS}

		t := template.Must(template.ParseFiles(
			"public/templates/docs.html",
			"public/templates/partials/meta.html",
			"public/templates/partials/navbar.html"))

		t.ExecuteTemplate(w, "docs", data)
		defer r.Body.Close()
	})
}

// Error404 : Not found error handler
func Error404(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"public/templates/404.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	t.ExecuteTemplate(w, "404", nil)
	defer r.Body.Close()
}
