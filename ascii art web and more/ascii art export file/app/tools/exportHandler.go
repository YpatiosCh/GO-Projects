package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ExportHandler handles the export of ASCII art in different formats (PDF, PNG, HTML, Text)
func ExportHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure that only POST requests are accepted
	if r.Method != http.MethodPost {
		RenderErrorTemplate(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// Parse the form data from the incoming request
	if err := r.ParseForm(); err != nil {
		RenderErrorTemplate(w, http.StatusBadRequest, "Error parsing form data")
		return
	}

	// Retrieve ASCII art and color information from the form data
	asciiArt := r.FormValue("ascii_art")
	if asciiArt == "" {
		RenderErrorTemplate(w, http.StatusBadRequest, "No ASCII art to export")
		return
	}

	// Retrieve the export format (e.g., JSON, XML, HTML, or default text)
	exportFormat := r.FormValue("format") // e.g., "JSON", "HTML"
	color := r.FormValue("color")
	backgroundColor := r.FormValue("background_color")

	// Switch statement to handle different export formats
	switch exportFormat {
	case "html":
		// Create HTML export
		htmlContent := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>ASCII Art Export</title>
				<style>
					body {
						color: %s;
						background-color: %s;
						font-family: monospace;
						white-space: pre-wrap; /* Ensures proper display of ASCII art */
					}
				</style>
			</head>
			<body>
			<pre>%s</pre>
			</body>
			</html>`, color, backgroundColor, asciiArt)

		// Set headers for HTML download
		filename := fmt.Sprintf("ascii_art_export_%d.html", time.Now().UnixNano())
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		w.Write([]byte(htmlContent)) // Write the HTML content to the response

	case "json":
		// Create a struct to hold the JSON data
		jsonData := struct {
			ASCIIArt        string `json:"ascii_art"`
			Color           string `json:"color"`
			BackgroundColor string `json:"background_color"`
		}{
			ASCIIArt:        asciiArt,
			Color:           color,
			BackgroundColor: backgroundColor,
		}

		// Marshal the struct to pretty-printed JSON
		prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			RenderErrorTemplate(w, http.StatusInternalServerError, "Error generating JSON")
			return
		}

		// Convert the ASCII art to a raw string with escaped newlines for better formatting
		prettyJSON = []byte(strings.ReplaceAll(string(prettyJSON), "\\n", "\n"))

		filename := fmt.Sprintf("ascii_art_export_%s.json", time.Now().Format("2006-01-02_15-04-05"))
		// Set headers for JSON download
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		w.Write(prettyJSON) // Write the JSON content to the response

	case "xml":
		// Create XML export
		xmlContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
		<ASCIIArtExport>
			<ASCIIArt><![CDATA[%s]]></ASCIIArt>
			<Color>%s</Color>
			<BackgroundColor>%s</BackgroundColor>
		</ASCIIArtExport>`, asciiArt, color, backgroundColor)

		// Set headers for XML download
		filename := fmt.Sprintf("ascii_art_export_%s.xml", time.Now().Format("2006-01-02_15-04-05"))
		w.Header().Set("Content-Type", "application/xml")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		w.Write([]byte(xmlContent)) // Write the XML content to the response

	default:
		// Default to text file export if no format is specified
		filename := fmt.Sprintf("ascii_art_export_%s.txt", time.Now().Format("2006-01-02_15-04-05"))
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		w.Write([]byte(asciiArt)) // Write the ASCII art as plain text to the response
	}
}
