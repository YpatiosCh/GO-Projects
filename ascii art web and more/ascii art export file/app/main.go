package main

import (
	"ascii-art-web-export-file/tools"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the 'static' directory (e.g., CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve files from the 'api' directory for API requests
	//http.Handle("static/api/", http.StripPrefix("static/api/", http.FileServer(http.Dir("api"))))
	// Serve API HTML (serve the api.html from the 'static/api' folder)
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/api/api.html") // Serve api.html file
	})

	// Handle the root path and form submission routes
	http.HandleFunc("/", tools.HomeHandler)                     // Handles GET requests to the home page
	http.HandleFunc("/ascii-art", tools.AsciiArtHandler)        // Handles POST requests for ASCII art generation
	http.HandleFunc("/api/ascii-art", tools.ApiAsciiArtHandler) // Handles POST requests for ASCII art via API
	http.HandleFunc("/export", tools.ExportHandler)             //Handles export requests for ASCII

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Starts the server on port 8080
}
