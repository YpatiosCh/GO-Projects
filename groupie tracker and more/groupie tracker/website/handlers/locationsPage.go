package handlers

import (
	"groupie-tracker/models"
	"groupie-tracker/utils"
	"log"
	"net/http"
)

// HandleLocation handles the "/location" endpoint
// Shows all concerts happening at a specific location
// Displays:
// - List of artists performing at this location
// - Concert dates for each artist
// Query parameter "q" is used for the location name
func HandleLocation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("q")
	log.Printf("Received location: %s", location)
	if location == "" {
		handleError(w, "Location not specified", http.StatusBadRequest)
		return
	}

	artists, datesLocations, err := utils.FetchLocationDetails(location)
	if err != nil {
		handleError(w, "Failed to fetch location data", http.StatusInternalServerError)
		return
	}

	if len(artists) == 0 {
		handleError(w, "No concerts found at this location", http.StatusNotFound)
		return
	}

	// Get dates from datesLocations map keys
	var dates []string
	for date := range datesLocations {
		dates = append(dates, date)
	}

	data := models.PageData{
		Title: "Concerts in " + location,
		Data: &models.ViewData{
			Location:       location,
			Artists:        artists,
			DatesLocations: datesLocations,
			Dates:          dates,
		},
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template error: %v", err)
		handleError(w, "Template error", http.StatusInternalServerError)
	}
}
