package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"groupie-tracker/models"
	"groupie-tracker/utils"
)

// init is called automatically when the package is initialized
// Here it initializes the search index before the server starts
func init() {
	if err := initSearchIndex(); err != nil {
		log.Fatal(err)
	}
}

// HandleSearch handles the "/api/search" endpoint
// Provides search functionality for:
// - Artist names
// - Band members
// - Locations
// - Concert dates
// Returns JSON response with up to 10 matching results
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	if query == "" {
		json.NewEncoder(w).Encode([]models.SearchResult{})
		return
	}

	// Track unique results using a map to avoid duplicates
	results := make([]models.SearchResult, 0)
	seen := make(map[string]bool)

	// Search through the index for matching results
	for indexKey, indexResults := range searchIndex {
		if strings.Contains(indexKey, query) {
			for _, result := range indexResults {
				uniqueKey := fmt.Sprintf("%s-%s-%d", result.Text, result.Type, result.ArtistID)
				if !seen[uniqueKey] {
					// Format location text before sending
					if result.Type == "location" || result.Type == "venue" {
						result.Text = utils.FormatLocation(result.Text)
					}
					results = append(results, result)
					seen[uniqueKey] = true
				}
			}
		}
	}

	// Limit results to top 10
	if len(results) > 10 {
		results = results[:10]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// initSearchIndex creates a search index for fast searching
// It indexes:
// - Artist names
// - Band members
// - Concert locations
// - Concert dates
// This is called once when the server starts
func initSearchIndex() error {
	artists, err := utils.FetchArtists()
	if err != nil {
		return err
	}

	searchIndex = make(map[string][]models.SearchResult)

	// First index artists and their members
	for _, artist := range artists {
		// Index artist name
		addToSearchIndex(strings.ToLower(artist.Name), models.SearchResult{
			Text:     artist.Name,
			Type:     "artist",
			ArtistID: artist.ID,
		})

		// Index members
		for _, member := range artist.Members {
			addToSearchIndex(strings.ToLower(member), models.SearchResult{
				Text:     member,
				Type:     "member of " + artist.Name,
				ArtistID: artist.ID,
			})

			// If member is also a solo artist, index them twice
			for _, a := range artists {
				if strings.EqualFold(member, a.Name) {
					addToSearchIndex(strings.ToLower(member), models.SearchResult{
						Text:     member,
						Type:     "artist",
						ArtistID: a.ID,
					})
				}
			}
		}

		// Index creation date
		addToSearchIndex(fmt.Sprintf("%d", artist.CreationDate), models.SearchResult{
			Text:     artist.Name,
			Type:     fmt.Sprintf("created in %d", artist.CreationDate),
			ArtistID: artist.ID,
		})

		// Index first album
		addToSearchIndex(artist.FirstAlbum, models.SearchResult{
			Text:     artist.Name,
			Type:     fmt.Sprintf("first album on %s", artist.FirstAlbum),
			ArtistID: artist.ID,
		})

		// Index locations
		relations, err := utils.FetchRelation(fmt.Sprint(artist.ID))
		if err != nil {
			continue
		}

		for location, dates := range relations.DatesLocations {

			// First, add the location itself
			addToSearchIndex(strings.ToLower(location), models.SearchResult{
				Text:         utils.FormatLocation(location),
				OriginalText: location,
				Type:         "location",
				ArtistID:     0,
			})

			// Then add it as a venue for this artist
			addToSearchIndex(strings.ToLower(location), models.SearchResult{
				Text:     utils.FormatLocation(location),
				Type:     "venue for " + artist.Name,
				ArtistID: artist.ID,
			})

			// Add concert dates
			for _, date := range dates {
				addToSearchIndex(strings.ToLower(date), models.SearchResult{
					Text:     date,
					Type:     "concert date",
					ArtistID: 0, // We use 0 because we want to link to date.html
				})
			}

		}
	}

	return nil
}

// addToSearchIndex adds a new entry to the search index
// It appends the result to the slice of results for the given key
func addToSearchIndex(key string, result models.SearchResult) {
	searchIndex[key] = append(searchIndex[key], result)
}
