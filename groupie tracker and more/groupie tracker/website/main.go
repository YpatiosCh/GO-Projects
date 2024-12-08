package main

import (
	"log"
	"net/http"

	"groupie-tracker/handlers"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artist/", handlers.HandleArtist)
	http.HandleFunc("/location", handlers.HandleLocation)
	http.HandleFunc("/date", handlers.HandleDate)
	http.HandleFunc("/api/search", handlers.HandleSearch)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
