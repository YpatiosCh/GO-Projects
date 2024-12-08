package models

// API response types
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// ViewData handles all view-related data
type ViewData struct {
	// Used in both location and date views to store the current location/date being viewed
	Location string
	Date     string

	// Used across all views (artist list, location concerts, date concerts)
	// Contains array of Artist objects with their details
	Artists []Artist

	// Used in artist view to show all locations and their concert dates
	// map[location] = []concertDates
	Relations map[string][]string

	// Used in location view to map concert dates to performing artists
	// map[date] = artistPerformingOnThatDate
	DatesLocations map[string]Artist

	// Used in location view to display list of all concert dates at this location
	Dates []string

	// Used in date view to display list of all locations with concerts on this date
	Locations []string

	// Used in date view to map locations to performing artists
	// map[location] = artistPerformingAtThatLocation
	LocationArtists map[string]Artist
}

// SearchResult represents a single search suggestion
type SearchResult struct {
	// Text is what's shown to the user in the search suggestions
	// Example: "Mexico City, Mexico" (formatted nicely for display)
	Text string `json:"text"`

	// OriginalText is the raw format used by the API
	// Example: "mexico-city_mexico"
	OriginalText string `json:"originalText"`

	// Type tells us what kind of result this is:
	// - "artist" (like "Queen")
	// - "member" (like "Freddie Mercury")
	// - "location" (like "London, UK")
	// - "venue" (same location but tied to specific artist)
	// - "concert date" (like "2024-01-01")
	Type string `json:"type"`

	// ArtistID is the unique identifier for the artist
	// Used to create links to artist pages
	ArtistID int `json:"artistId"`

	// Context provides additional information:
	// - For members: the band they're in
	// - For venues: the artist performing there
	// omitempty means this field is optional in JSON
	Context string `json:"context,omitempty"`
}

// PageData is used for rendering HTML templates
type PageData struct {
	// Title appears in browser tab and page header
	// Example: "Concerts in London" or "Queen"
	Title string

	// IsHome is true only on the main page
	// Used to show/hide certain elements (like search bar)
	IsHome bool

	// Data contains all the information needed for the page
	// Pointer to ViewData struct which has artists, locations, etc.
	Data *ViewData

	// Error message to display if something goes wrong
	// Example: "Artist not found" or "Invalid location"
	Error string
}
