package handlers

import (
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"log"
	"net/http"
)

// HandleDate handles the "/date" endpoint
// Shows all concerts happening on a specific date
// Displays:
// - List of all artists performing on this date
// Query parameter "q" is used for the date
func HandleDate(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("q")
	if date == "" {
		handleError(w, "Date not specified", http.StatusBadRequest)
		return
	}

	artists, locations, locationArtists, err := utils.FetchDateDetails(date)
	if err != nil {
		handleError(w, "Failed to fetch date data", http.StatusInternalServerError)
		return
	}

	data := models.PageData{
		Title: "Concerts on " + date,
		Data: &models.ViewData{
			Date:            date,
			Artists:         artists,
			Locations:       locations,
			LocationArtists: locationArtists,
		},
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template error: %v", err)
		handleError(w, "Template error", http.StatusInternalServerError)
	}
}
