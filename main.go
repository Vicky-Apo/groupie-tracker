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
