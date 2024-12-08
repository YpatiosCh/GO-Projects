package main

import (
	"ascii-art-web-export-file/tools"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHomeHandler tests the home page handler for GET requests.
func TestHomeHandler(t *testing.T) {
	// Create a GET request for the root ("/") endpoint
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err) // Fail the test if the request cannot be created
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tools.HomeHandler)

	// Call the HomeHandler function
	handler.ServeHTTP(rr, req)

	// Check that the response status code is 200 (OK)
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, status)
	}

	// Get the response body as a string
	body := rr.Body.String()

	// Check for specific HTML elements in the response to verify the page content
	if !strings.Contains(body, "<title>ASCII Art Generator</title>") {
		t.Errorf("Expected title not found in homeHandler response")
	}
	if !strings.Contains(body, `form action="/ascii-art" method="post"`) {
		t.Errorf("Expected form action not found in homeHandler response")
	}
	if !strings.Contains(body, `<textarea id="text" name="text" required>`) {
		t.Errorf("Expected textarea for text input not found in homeHandler response")
	}
	if !strings.Contains(body, `input type="radio" id="standard" name="banner" value="standard" required`) ||
		!strings.Contains(body, `input type="radio" id="shadow" name="banner" value="shadow"`) ||
		!strings.Contains(body, `input type="radio" id="thinkertoy" name="banner" value="thinkertoy"`) {
		t.Errorf("Expected banner radio buttons not found in homeHandler response")
	}
	if !strings.Contains(body, `<button type="submit">Generate</button>`) {
		t.Errorf("Expected submit button not found in homeHandler response")
	}
}

// TestAsciiArtHandler tests the ASCII art generation handler with valid input.
func TestAsciiArtHandler(t *testing.T) {
	// Simulate form data for ASCII art generation
	form := "text=Hello&banner=standard"
	req, err := http.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tools.AsciiArtHandler)

	// Call the AsciiArtHandler function
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("asciiArtHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify the response contains ASCII art output
	if !strings.Contains(rr.Body.String(), "H") {
		t.Errorf("asciiArtHandler did not generate ASCII art correctly")
	}
}

// TestAsciiArtHandlerBadRequest tests ASCII art generation handler with missing form data.
func TestAsciiArtHandlerBadRequest(t *testing.T) {
	// Create a POST request without form data
	req, err := http.NewRequest(http.MethodPost, "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tools.AsciiArtHandler)

	// Call the AsciiArtHandler function
	handler.ServeHTTP(rr, req)

	// Check the response status code for a bad request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("asciiArtHandler returned wrong status code for bad request: got %v want %v", status, http.StatusBadRequest)
	}
}

// TestApiAsciiArtHandler tests the API ASCII art handler with JSON input.
func TestApiAsciiArtHandler(t *testing.T) {
	// Simulate a JSON request with text and banner type
	requestBody := `{"text": "Hello", "banner": "standard"}`
	req, err := http.NewRequest(http.MethodPost, "/api/ascii-art", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tools.ApiAsciiArtHandler)
	handler.ServeHTTP(rr, req)

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rr.Code)
	}

	// Decode the JSON response into a struct
	var response struct {
		AsciiArt []string `json:"ascii_art"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response JSON: %v", err)
	}

	// Log the ASCII art for debugging purposes
	for i, line := range response.AsciiArt {
		t.Logf("ASCII Art Line %d: %s", i, line)
	}

	// Verify the ASCII art output starts with the expected pattern
	expectedStart := " _    _" // Adjust based on actual output
	if len(response.AsciiArt) == 0 || !strings.HasPrefix(response.AsciiArt[0], expectedStart) {
		t.Errorf("apiAsciiArtHandler did not generate ASCII art correctly, expected start with: %v", expectedStart)
	}
}

// TestExportHandler tests the ExportHandler for different export formats.
func TestExportHandler(t *testing.T) {
	// Define test cases for export functionality
	tests := []struct {
		name           string
		method         string
		formData       string
		expectedStatus int
		expectedType   string
	}{
		{"Export HTML", http.MethodPost, "ascii_art=Test+Art&format=html&color=#000000&background_color=#FFFFFF", http.StatusOK, "text/html"},
		{"Export JSON", http.MethodPost, "ascii_art=Line1%0ALine2%0ALine3&format=json&color=#123456&background_color=#654321", http.StatusOK, "application/json"},
		{"Export XML", http.MethodPost, "ascii_art=Art%20In%20XML&format=xml&color=#111111&background_color=#222222", http.StatusOK, "application/xml"},
		{"Export Text (default)", http.MethodPost, "ascii_art=Default+Text", http.StatusOK, "text/plain"},
		{"Invalid Method", http.MethodGet, "", http.StatusMethodNotAllowed, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request for each test case
			req := httptest.NewRequest(tt.method, "/export", strings.NewReader(tt.formData))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tools.ExportHandler)
			handler.ServeHTTP(rr, req)

			// Verify the status code
			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			// Check the response content type if applicable
			if tt.expectedType != "" {
				contentType := rr.Header().Get("Content-Type")
				if !strings.HasPrefix(contentType, tt.expectedType) {
					t.Errorf("expected content type %q, got %q", tt.expectedType, contentType)
				}
			}

			// Log the response body for manual verification during testing
			t.Logf("Response Body: %s", rr.Body.String())
		})
	}
}
