package Views

import (
	"html/template"
	"net/http"

	"waifu.pics/util"
)

type Front struct {
	Endpoints []string
}

type Multi struct {
	Endpoint string
	Config   util.Config
}

type grid struct {
	URL      string
	Endpoint string
}

// Grid : This is the grid page initializer for every endpoint
func (multi Multi) Grid(w http.ResponseWriter, r *http.Request) {
	p := grid{URL: multi.Config.URL, Endpoint: multi.Endpoint}
	// Setting up all templates
	t := template.Must(template.ParseFiles(
		"public/templates/grid.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	t.ExecuteTemplate(w, "grid", p)
	defer r.Body.Close()
}

// Docs : This is the docs page
func (front *Front) Docs(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"public/templates/docs.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	t.ExecuteTemplate(w, "docs", front.Endpoints)
	defer r.Body.Close()
}

// UploadFront : This is the upload page
func (front *Front) UploadFront(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"public/templates/upload.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	t.ExecuteTemplate(w, "upload", front.Endpoints)
	defer r.Body.Close()
}

// AdminLogin : Login page for admins
func AdminLogin(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"public/templates/admin/login.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	t.ExecuteTemplate(w, "adminlogin", nil)
	defer r.Body.Close()
}

// AdminDash : Admin dashboard page
func AdminDash(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"public/templates/admin/dash.html",
		"public/templates/partials/meta.html",
		"public/templates/partials/navbar.html"))

	t.ExecuteTemplate(w, "admindash", nil)
	defer r.Body.Close()
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
