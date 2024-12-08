package handlers

import (
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// HandleArtist handles the "/artist/{id}" endpoint
// Displays detailed information about a specific artist including:
// - Basic info (name, image, etc.)
// - Concert locations
// - Concert dates
// - Relations between dates and locations
func HandleArtist(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/artist/")
	if id == "" {
		handleError(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}
	idInt, _ := strconv.Atoi(id)
	if idInt > 52 {
		id = "52"
	}

	artist, relations, err := utils.FetchArtistDetailsAndRelations(id)
	if err != nil {
		handleError(w, "Failed to fetch artist data", http.StatusInternalServerError)
		return
	}

	data := models.PageData{
		Title: artist.Name,
		Data: &models.ViewData{
			Artists:   []models.Artist{artist},
			Relations: relations.DatesLocations,
		},
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template error: %v", err)
		handleError(w, "Template error", http.StatusInternalServerError)
	}
}
