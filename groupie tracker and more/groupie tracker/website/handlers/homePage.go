package handlers

import (
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"html/template"
	"log"
	"net/http"
)

// Global variables for template parsing and search functionality
var templates = template.Must(template.New("").Funcs(template.FuncMap{
	"FormatLocation": utils.FormatLocation,
}).ParseGlob("templates/*.html"))
var searchIndex map[string][]models.SearchResult

// HandleHome handles the root ("/") endpoint
// It displays the main page with all artists in a grid layout
// Returns 404 if path is not exactly "/"
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, "Page Not Found", http.StatusNotFound)
		return
	}

	artists, err := utils.FetchArtists()
	if err != nil {
		handleError(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	data := models.PageData{
		IsHome: true,
		Data:   &models.ViewData{Artists: artists},
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template error: %v", err)
		handleError(w, "Template error", http.StatusInternalServerError)
	}
}
