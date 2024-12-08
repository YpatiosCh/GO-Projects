package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// Parse the HTML templates used for rendering
var templates = template.Must(template.ParseFiles("templates/index.html", "templates/result.html", "templates/errors.html"))

func main() {
    // Serve static files from the 'static' directory (e.g., CSS, JS, images)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))	
    
    // Serve files from the 'api' directory for API requests
    http.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir("api"))))

    // Handle the root path and form submission routes
    http.HandleFunc("/", homeHandler)              // Handles GET requests to the home page
    http.HandleFunc("/ascii-art", asciiArtHandler) // Handles POST requests for ASCII art generation
	http.HandleFunc("/api/ascii-art", apiAsciiArtHandler) // Handles POST requests for ASCII art via API

    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Starts the server on port 8080
}

// Home handler for displaying the form
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the URL path is incorrect (not "/"), return 404 error
    if r.URL.Path != "/" {
        renderErrorTemplate(w, http.StatusNotFound, "Page Not Found")
        return
    }

    // Check if the HTTP method is not GET (only allow GET requests for the home page)
    if r.Method != http.MethodGet {
        renderErrorTemplate(w, http.StatusMethodNotAllowed, "Method Not Allowed")
        return
    }

    // Render the home page (index.html)
    if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
        renderErrorTemplate(w, http.StatusInternalServerError, "Error rendering template")
    }
}

// Handler for generating ASCII art from form input (POST request)
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
    // If the request method is not POST, return a 405 error
    if r.Method != http.MethodPost {
        renderErrorTemplate(w, http.StatusMethodNotAllowed, "Method Not Allowed")
        return
    }

    // Get the 'text' and 'banner' values from the form submission
    text := r.FormValue("text")
    banner := r.FormValue("banner")

    // If either 'text' or 'banner' is missing, return a 400 error
    if text == "" || banner == "" {
        renderErrorTemplate(w, http.StatusBadRequest, "Missing text or banner")
        return
    }

    // Construct the file path to the chosen banner
    fontFile := fmt.Sprintf("banners/%s.txt", banner)

    // Call the function to generate ASCII art
    result, err := printAsciiArt(text, fontFile)
    if err != nil {
        renderErrorTemplate(w, http.StatusInternalServerError, fmt.Sprintf("Error generating ASCII art: %v", err))
        return
    }

    // Prepare the data to pass to the result template
    data := struct {
        AsciiArt string
    }{
        AsciiArt: strings.Join(result, "\n"), // Join the ASCII art lines with newlines
    }

    // Render the result page (result.html) with the generated ASCII art
    if err := templates.ExecuteTemplate(w, "result.html", data); err != nil {
        renderErrorTemplate(w, http.StatusInternalServerError, "Error rendering result")
    }
}

// Handler for generating ASCII art via API and returning the result in JSON format
func apiAsciiArtHandler(w http.ResponseWriter, r *http.Request) {
    // If the request method is not POST, return a 405 error
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed) // 405 Method Not Allowed
        return
    }

    // Create a struct to hold the incoming JSON data
    var requestData struct {
        Text   string `json:"text"`
        Banner string `json:"banner"`
    }

    // Parse the JSON request body into the struct
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest) // 400 Bad Request
        return
    }

    // Check if the required fields (text or banner) are missing
    if requestData.Text == "" || requestData.Banner == "" {
        http.Error(w, "Missing text or banner", http.StatusBadRequest) // 400 Bad Request
        return
    }

    // Construct the file path to the chosen banner
    fontFile := fmt.Sprintf("banners/%s.txt", requestData.Banner)

    // Call the function to generate ASCII art
    result, err := printAsciiArt(requestData.Text, fontFile)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error generating ASCII art: %v", err), http.StatusInternalServerError) // 500 Internal Server Error
        return
    }

    // Prepare the response data (ASCII art as an array of strings)
    response := struct {
        AsciiArt []string `json:"ascii_art"`
    }{
        AsciiArt: result,
    }

    // Set the response content type to JSON and encode the response
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // 500 Internal Server Error
    }
}

// Print ASCII art for the given text and banner
func printAsciiArt(word string, filename string) ([]string, error) {
    // Open the font file for the selected banner
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("font not available. You can choose between: standard, shadow, thinkertoy")
    }
    defer file.Close()

    // Split the input text into words (if there are multiple words)
    words := strings.Split(strings.ReplaceAll(word, "\r\n", "\n"), "\n")
    var allLines []string

    // Loop through each word and generate ASCII art for it
    for wordIndex, currentWord := range words {
        if currentWord == "" {
            if wordIndex < len(words)-1 {
                allLines = append(allLines, "")
            }
            continue
        }

        lines := make([]string, 8) // Each letter has 8 lines in the ASCII art

        // Loop through each letter in the current word and generate ASCII art
        for _, letter := range currentWord {
            art, err := getAsciiArtForLetter(file, letter)
            if err != nil {
                return nil, fmt.Errorf("error reading ascii art for letter %c: %v", letter, err)
            }
            // Append the ASCII art for each letter to the lines
            for j := 0; j < 8; j++ {
                lines[j] += art[j]
            }
        }

        allLines = append(allLines, lines...)
        if wordIndex < len(words)-1 {
            allLines = append(allLines, "")
        }
    }

    return allLines, nil
}

// Get ASCII art for a specific letter from the font file
func getAsciiArtForLetter(file *os.File, letter rune) ([]string, error) {
    // Only process printable characters (ASCII values between 32 and 126)
    if letter < 32 || letter > 126 {
        return nil, fmt.Errorf("character %c is not printable", letter)
    }

    line := int(letter)
    lineFromText := (line-32)*9 + 1

    _, err := file.Seek(0, 0)
    if err != nil {
        return nil, err
    }

    scanner := bufio.NewScanner(file)

    // Skip to the correct starting line for the letter in the font file
    for i := 0; i < lineFromText; i++ {
        if !scanner.Scan() {
            return nil, fmt.Errorf("failed to find line %d for letter %c", lineFromText, letter)
        }
    }

    // Read the 8 lines of ASCII art for the character
    var asciiArt []string
    for i := 0; i < 8; i++ {
        if scanner.Scan() {
            asciiArt = append(asciiArt, scanner.Text())
        } else {
            return nil, fmt.Errorf("failed to read full ascii art for letter %c (expected 8 lines, got %d)", letter, i)
        }
    }

    return asciiArt, nil
}

// renderErrorTemplate renders the errors.html template with a status code and message
func renderErrorTemplate(w http.ResponseWriter, statusCode int, message string) {
    // Set the status code in the HTTP response
    w.WriteHeader(statusCode)

    // Define the data to pass to the template (status code and message)
    data := struct {
        StatusCode int
        Message    string
    }{
        StatusCode: statusCode,
        Message:    message,
    }

    // Render the error page using the errors.html template
    if err := templates.ExecuteTemplate(w, "errors.html", data); err != nil {
        http.Error(w, "Error rendering error page", http.StatusInternalServerError) // Handle template rendering errors
    }
}
