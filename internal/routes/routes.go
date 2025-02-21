package routes

import (
	"html/template"
	"net/http"

	"groupie-tracker/internal/handlers"
)

// render404 renders a custom 404 page.
func render404(tpl *template.Template, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	if err := tpl.ExecuteTemplate(w, "404.html", nil); err != nil {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
	}
}

// NewRouter sets up the mux with all the necessary routes.
func NewRouter(tpl *template.Template) *http.ServeMux {
	mux := http.NewServeMux()

	// Home route with exact path checking.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			render404(tpl, w)
			return
		}
		handlers.HomeHandler(tpl)(w, r)
	})

	// Artist detail route.
	mux.HandleFunc("/artist/", handlers.DetailHandler(tpl))

	// API Route for fetching artists
	mux.HandleFunc("/api/artists", handlers.GetArtists)

	// Trigger event route.
	mux.HandleFunc("/trigger-event", handlers.TriggerEventHandler)

	// Serve static files.
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
