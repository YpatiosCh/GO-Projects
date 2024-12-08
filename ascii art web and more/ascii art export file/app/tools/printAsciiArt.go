package tools

import (
	"os"
	"fmt"
	"strings"
)

// PrintAsciiArt generates ASCII art for the given text using a specific banner font.
// It splits the input text into words, and for each word, it retrieves the ASCII art for each letter 
// and combines them to form the complete ASCII art representation of the text.
//
// Parameters:
// - word: The text that needs to be converted to ASCII art (can include multiple words).
// - filename: The name of the font file containing the ASCII art patterns for the banner style.
//
// Returns:
// - A slice of strings, where each string is a line of ASCII art for the entire input text.
// - An error if there is an issue opening the font file or generating the ASCII art.
func PrintAsciiArt(word string, filename string) ([]string, error) {
	// Open the font file for the selected banner
	file, err := os.Open(filename)
	if err != nil {
		// Return an error if the font file cannot be opened
		return nil, fmt.Errorf("font not available. You can choose between: standard, shadow, thinkertoy")
	}
	// Ensure the file is closed after processing
	defer file.Close()

	// Split the input text into words (handles newline characters in input)
	words := strings.Split(strings.ReplaceAll(word, "\r\n", "\n"), "\n")
	var allLines []string

	// Loop through each word in the input text
	for wordIndex, currentWord := range words {
		// If the word is empty (e.g., extra line breaks), skip it
		if currentWord == "" {
			// If not the last word, add an empty line to maintain formatting
			if wordIndex < len(words)-1 {
				allLines = append(allLines, "")
			}
			continue
		}

		// Initialize an array to hold the ASCII art for the current word (each word has 8 lines)
		lines := make([]string, 8)

		// Loop through each letter in the current word to generate its ASCII art
		for _, letter := range currentWord {
			// Get the ASCII art for the letter using the provided font file
			art, err := GetAsciiArtForLetter(file, letter)
			if err != nil {
				// Return an error if there's a problem getting the ASCII art for the letter
				return nil, fmt.Errorf("error reading ascii art for letter %c: %v", letter, err)
			}
			// Append the ASCII art for each letter to the corresponding lines
			for j := 0; j < 8; j++ {
				lines[j] += art[j]
			}
		}

		// Append the lines for the current word to the final output
		allLines = append(allLines, lines...)
		// Add a blank line between words if not the last word
		if wordIndex < len(words)-1 {
			allLines = append(allLines, "")
		}
	}

	// Return the complete ASCII art for the entire input text
	return allLines, nil
}
