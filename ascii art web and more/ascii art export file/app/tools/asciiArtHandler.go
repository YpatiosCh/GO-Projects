package tools

import (
	"fmt"
	"net/http"
	"strings"
)

// AsciiArtHandler handles the ASCII art generation request
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST, otherwise return a "Method Not Allowed" error
	if r.Method != http.MethodPost {
		RenderErrorTemplate(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Retrieve the form values: 'text', 'banner', and 'color'
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	color := SanitizeColor(r.FormValue("color")) // Sanitize the color input to prevent issues

	// If 'text' or 'banner' are missing, return an error indicating the missing fields
	if text == "" || banner == "" {
		RenderErrorTemplate(w, http.StatusBadRequest, "Missing text or banner")
		return
	}

	// Construct the file path for the banner font (e.g., "banners/standard.txt")
	fontFile := fmt.Sprintf("banners/%s.txt", banner)
	
	// Call the PrintAsciiArt function to generate the ASCII art based on the input text and font file
	result, err := PrintAsciiArt(text, fontFile)
	if err != nil {
		// If there is an error generating the ASCII art, render an error template
		RenderErrorTemplate(w, http.StatusInternalServerError, fmt.Sprintf("Error generating ASCII art: %v", err))
		return
	}

	// Calculate the background color based on the text color for contrast
	backgroundColor := GetContrastingBackground(color)

	// Create a data struct to pass the ASCII art, text color, and background color to the template
	data := struct {
		AsciiArt        string // The generated ASCII art as a string
		Color           string // The selected text color
		BackgroundColor string // The contrasting background color
	}{
		AsciiArt:        strings.Join(result, "\n"), // Join the result slice into a single string
		Color:           color,                        // The color for the text
		BackgroundColor: backgroundColor,             // The contrasting background color
	}

	// Execute the "result.html" template, passing the data struct as context
	if err := templates.ExecuteTemplate(w, "result.html", data); err != nil {
		// If there is an error rendering the template, render an error template
		RenderErrorTemplate(w, http.StatusInternalServerError, "Error rendering result")
	}
}
