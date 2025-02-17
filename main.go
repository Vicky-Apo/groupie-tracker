package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"groupie-tracker/internal/data"
	"groupie-tracker/internal/handlers"
)

// ReplaceSpaces replaces spaces with dashes in artist names for clean URLs
func ReplaceSpaces(name string) string {
	return strings.ReplaceAll(name, " ", "-")
}

// render404 renders the custom 404 page using the provided template.
func render404(tpl *template.Template, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	if err := tpl.ExecuteTemplate(w, "404.html", nil); err != nil {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
	}
}

func main() {
	// Load data from API
	if err := data.LoadData(); err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	// Define custom template functions
	funcMap := template.FuncMap{
		"replaceSpaces": ReplaceSpaces,
	}

	// Parse all templates and apply the function map
	tmplPattern := filepath.Join("templates", "*.html")
	tpl := template.Must(template.New("").Funcs(funcMap).ParseGlob(tmplPattern))

	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Register the home handler with an exact path check.
	// If the URL path isn't exactly "/", we render the custom 404 page.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			render404(tpl, w)
			return
		}
		handlers.HomeHandler(tpl)(w, r)
	})

	// Register other handlers.
	mux.HandleFunc("/artist/", handlers.DetailHandler(tpl))
	mux.HandleFunc("/trigger-event", handlers.TriggerEventHandler)

	// Serve static files from the "static" directory.
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server.
	addr := ":8080"
	fmt.Printf("Server is running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
