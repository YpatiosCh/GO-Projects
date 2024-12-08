package tools

import (
	"os"
	"fmt"
	"bufio"
)

// GetAsciiArtForLetter retrieves the ASCII art representation of a given letter from a font file.
// The letter should be a printable ASCII character (between 32 and 126).
// The font file is expected to contain ASCII art for characters starting from ASCII 32 (space) and upwards.
func GetAsciiArtForLetter(file *os.File, letter rune) ([]string, error) {
	// Check if the character is a printable ASCII character (between 32 and 126)
	if letter < 32 || letter > 126 {
		// If the character is not printable, return an error
		return nil, fmt.Errorf("character %c is not printable", letter)
	}

	// Convert the letter to an integer, used for finding the correct line in the font file
	line := int(letter)
	// Calculate the line number where the ASCII art for the letter starts in the font file.
	// Each letter takes 9 lines, and the first letter (ASCII 32) starts at line 1.
	lineFromText := (line-32)*9 + 1

	// Go to the beginning of the file to start scanning from the top
	_, err := file.Seek(0, 0)
	if err != nil {
		// If there is an error seeking the file, return it
		return nil, err
	}

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Skip over lines until we reach the starting line for the letter
	for i := 0; i < lineFromText; i++ {
		if !scanner.Scan() {
			// If the scanner reaches the end of the file before the expected line, return an error
			return nil, fmt.Errorf("failed to find line %d for letter %c", lineFromText, letter)
		}
	}

	// Create a slice to store the 8 lines of ASCII art for the letter
	var asciiArt []string
	// Read the next 8 lines (representing the ASCII art for the letter)
	for i := 0; i < 8; i++ {
		if scanner.Scan() {
			// Append each line of the ASCII art to the slice
			asciiArt = append(asciiArt, scanner.Text())
		} else {
			// If fewer than 8 lines are found, return an error
			return nil, fmt.Errorf("failed to read full ascii art for letter %c (expected 8 lines, got %d)", letter, i)
		}
	}

	// Return the ASCII art for the letter
	return asciiArt, nil
}
