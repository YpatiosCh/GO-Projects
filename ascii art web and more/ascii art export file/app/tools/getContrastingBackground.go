package tools

import "fmt"

// GetContrastingBackground determines a contrasting background color based on the input text color.
// It calculates the brightness of the provided color and returns either a light or dark background color 
// for better readability of the text.
//
// Parameters:
// - color: A hex color code string (e.g., "#RRGGBB") representing the text color.
//
// Returns:
// - A hex color code for the background color that contrasts with the given text color.
func GetContrastingBackground(color string) string {
    // Convert the provided hex color string to RGB components.
    var r, g, b int
    // The fmt.Sscanf function reads the RGB values from the hex string (e.g., "#ff5733" -> r=255, g=87, b=51)
    fmt.Sscanf(color, "#%02x%02x%02x", &r, &g, &b)

    // Calculate the brightness of the color using the luminance formula.
    // The luminance formula is a weighted sum that accounts for the human eye's sensitivity to different colors.
    // Here, red contributes 29.9%, green 58.7%, and blue 11.4% to the overall brightness.
    brightness := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)

    // If the brightness is below a threshold (128), return a light background (white).
    // A lower brightness means the color is darker, so a light background improves contrast.
    if brightness < 128 {
        return "#FFFFFF" // Light background (white) for dark text
    }
    
    // If the brightness is above or equal to 128, return a dark background (black).
    // A higher brightness means the color is lighter, so a dark background improves contrast.
    return "#000000" // Dark background (black) for light text
}
