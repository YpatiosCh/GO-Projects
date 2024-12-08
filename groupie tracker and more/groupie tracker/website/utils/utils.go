package utils

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"log"
	"net/http"
	"strings"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

// Generic function to fetch and unmarshal data
func fetchAndUnmarshal(endpoint string, target interface{}) error {
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

func FetchArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := fetchAndUnmarshal("/artists", &artists)
	return artists, err
}

func FetchArtistDetailsAndRelations(id string) (models.Artist, models.Relation, error) {
	var artist models.Artist
	var relation models.Relation

	if err := fetchAndUnmarshal("/artists/"+id, &artist); err != nil {
		return artist, relation, err
	}

	if err := fetchAndUnmarshal("/relation/"+id, &relation); err != nil {
		return artist, relation, err
	}

	return artist, relation, nil
}

// Add these helper functions in utils.go
func FormatToAPILocation(location string) string {
	// Remove any spaces around commas
	location = strings.ReplaceAll(location, ", ", ",")
	// Replace spaces with hyphens
	location = strings.ReplaceAll(location, " ", "_")
	// Replace commas with underscores
	location = strings.ReplaceAll(location, ",", "-")
	// Convert to lowercase
	return strings.ToLower(location)
}

// Then update FetchLocationDetails
func FetchLocationDetails(location string) ([]models.Artist, map[string]models.Artist, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, nil, err
	}

	// Convert the formatted location back to API format
	apiLocation := FormatToAPILocation(location)
	log.Print(apiLocation)

	var locationArtists []models.Artist
	datesLocations := make(map[string]models.Artist)

	for _, artist := range artists {
		var relation models.Relation
		if err := fetchAndUnmarshal("/relation/"+fmt.Sprint(artist.ID), &relation); err != nil {
			continue
		}

		if dates, exists := relation.DatesLocations[apiLocation]; exists {
			locationArtists = append(locationArtists, artist)
			for _, date := range dates {
				datesLocations[date] = artist
			}
		}
	}

	return locationArtists, datesLocations, nil
}

func FetchDateDetails(date string) ([]models.Artist, []string, map[string]models.Artist, error) {
	artists, err := FetchArtists()
	if err != nil {
		return nil, nil, nil, err
	}

	var dateArtists []models.Artist
	locationSet := make(map[string]bool) // To track unique locations
	var locations []string
	locationArtists := make(map[string]models.Artist)

	for _, artist := range artists {
		var relation models.Relation
		if err := fetchAndUnmarshal("/relation/"+fmt.Sprint(artist.ID), &relation); err != nil {
			continue
		}

		// Check each location and its dates
		for location, dates := range relation.DatesLocations {
			for _, d := range dates {
				if d == date {
					dateArtists = append(dateArtists, artist)
					if !locationSet[location] {
						locations = append(locations, location)
						locationSet[location] = true
					}
					locationArtists[location] = artist
					break
				}
			}
		}
	}

	return dateArtists, locations, locationArtists, nil
}

func FetchRelation(id string) (models.Relation, error) {
	var relation models.Relation
	err := fetchAndUnmarshal("/relation/"+id, &relation)
	return relation, err
}

func FormatLocation(location string) string {
	// Replace hyphens and underscores
	location = strings.ReplaceAll(location, "-", ", ")
	location = strings.ReplaceAll(location, "_", " ")

	// Split into words
	words := strings.Split(location, " ")

	wordsCap := Capitalize(words)

	return strings.Join(wordsCap, " ")
}

func Capitalize(words []string) []string {
	for i, word := range words {
		if word == "usa" {
			words[i] = "USA"
		} else if word == "uk" {
			words[i] = "UK"
		} else {
			words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
		}
	}

	return words
}
