package handlers

import (
	"groupie-tracker/models"
	"net/http"
)

// handleError displays an error page to the user
// It sets the appropriate HTTP status code and shows an error message
func handleError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	data := models.PageData{
		Title: "Error",
		Error: message,
	}
	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		http.Error(w, message, code)
	}
}
