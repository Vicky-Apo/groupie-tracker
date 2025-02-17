package handlers

import (
	"groupie-tracker/internal/data"
	"html/template"
	"net/http"
	"sort"
)

// HomeHandler handles the "/" route
func HomeHandler(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(data.AllArtists) == 0 {
			http.Error(w, "No artist data available", http.StatusInternalServerError)
			return
		}

		// Sort artists by name
		sortedArtists := make([]data.Artist, len(data.AllArtists))
		copy(sortedArtists, data.AllArtists)
		sort.Slice(sortedArtists, func(i, j int) bool {
			return sortedArtists[i].Name < sortedArtists[j].Name
		})

		// Execute the homepage template
		err := tpl.ExecuteTemplate(w, "home.html", sortedArtists)
		if err != nil {
			http.Error(w, "Internal Server Error while rendering index", http.StatusInternalServerError)
			return
		}
	}
}
