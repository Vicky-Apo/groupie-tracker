package handlers

import (
	"encoding/json"
	"groupie-tracker/internal/data"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

// HomeHandler handles the "/" route and renders the home page.
func HomeHandler(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if r.URL.Path != "/" {
			handler404(tpl, w)
			return
		}

		if len(data.AllArtists) == 0 {
			log.Println("ERROR: No artist data available")
			http.Error(w, "No artist data available", http.StatusInternalServerError)
			return
		}

		// Sort artists by name, case insensitive
		sortedArtists := make([]data.Artist, len(data.AllArtists))
		copy(sortedArtists, data.AllArtists)
		sort.Slice(sortedArtists, func(i, j int) bool {
			return strings.ToLower(sortedArtists[i].Name) < strings.ToLower(sortedArtists[j].Name)
		})

		// Render the homepage template
		err := tpl.ExecuteTemplate(w, "home.html", sortedArtists)
		if err != nil {
			log.Println("ERROR rendering template:", err)
			http.Error(w, "Internal Server Error while rendering index", http.StatusInternalServerError)
		}
	}
}

// render404 renders a custom 404 page.
func handler404(tpl *template.Template, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	if err := tpl.ExecuteTemplate(w, "404.html", nil); err != nil {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
	}
}

// GetArtists handles API requests for fetching all artists
func GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(data.AllArtists) == 0 {
		log.Println("ERROR: No artist data available")
		http.Error(w, "No artist data available", http.StatusInternalServerError)
		return
	}

	// Sort artists by name, case insensitive
	sortedArtists := make([]data.Artist, len(data.AllArtists))
	copy(sortedArtists, data.AllArtists)
	sort.Slice(sortedArtists, func(i, j int) bool {
		return strings.ToLower(sortedArtists[i].Name) < strings.ToLower(sortedArtists[j].Name)
	})

	if err := json.NewEncoder(w).Encode(sortedArtists); err != nil {
		log.Println("ERROR encoding JSON:", err)
		http.Error(w, "Internal Server Error while encoding JSON", http.StatusInternalServerError)
	}
}
