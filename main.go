package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"groupie-tracker/internal/data"
	"groupie-tracker/internal/handlers"
)

func main() {
	// Load data from API
	if err := data.LoadData(); err != nil {
		log.Fatalf("Error loading data: %v", err)
	}
	// Parse all templates in templates/*.html
	tmplPattern := filepath.Join("templates", "*.html")
	tpl, err := template.ParseGlob(tmplPattern)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// Register handlers
	http.HandleFunc("/", handlers.HomeHandler(tpl))
	http.HandleFunc("/artist/", handlers.DetailHandler(tpl))
	http.HandleFunc("/trigger-event", handlers.TriggerEventHandler)

	// Serve static files from web/static
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start server
	addr := ":8080"
	fmt.Printf("Server is running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
