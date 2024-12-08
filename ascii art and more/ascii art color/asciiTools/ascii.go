package asciiTools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to print ASCII art for a word or a whole index of a []string with two optional colors
func printAsciiArt(word string, filename string, primaryColors []string, secondaryColors []string, colorFlags []bool) ([]string, error) {
	// Open the ASCII art file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("font not available. You can choose between: <standard> <shadow> <thinkertoy>")
	}
	defer file.Close()

	lines := make([]string, 8) // Each letter has 8 lines in the ASCII art

	// Iterate over each letter in the word
	for i, letter := range word {
		art, err := GetAsciiArtForLetter(file, letter)
		if err != nil {
			return nil, fmt.Errorf("error reading ascii art for letter %c: %v", letter, err)
		}

		// Choose color based on the current index
		currentColor := Reset
		if colorFlags[i] {
			if len(secondaryColors) > 0 && i%2 != 0 {
				currentColor = secondaryColors[i%len(secondaryColors)]
			} else {
				currentColor = primaryColors[i%len(primaryColors)]
			}
		}

		// Append each line of the letter to the corresponding line in 'lines'
		for j := 0; j < 8; j++ {
			lines[j] += fmt.Sprintf("%s%s%s", currentColor, art[j], Reset)
		}
	}

	return lines, nil
}

// GetAsciiArtForLetter reads ASCII art for a specific character from the file
func GetAsciiArtForLetter(file *os.File, letter rune) ([]string, error) {
	line := int(letter)
	lineFromText := (line-32)*9 + 1 // Adjusted formula to find the starting line

	// Move to the appropriate line in the file
	_, err := file.Seek(0, 0) // Start from the beginning of the file
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	// Skip lines until we reach the starting line for the letter
	for i := 0; i < lineFromText; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("failed to find line %d for letter %c", lineFromText, letter)
		}
	}

	// Read the next 8 lines to get the ASCII art for the letter
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

// PrintFullAscii prints the ASCII art for the input string to both the terminal and a file (if outputFile is specified)
func PrintFullAscii(substring, fullInput, colorFlag, color2Flag, font, outputFile string) {
	// Split the input on '\n' to handle multi-line input
	parts := strings.Split(fullInput, "\\n")
	primaryColors := determineColors(&colorFlag)
	secondaryColors := []string{}
	if color2Flag != "" {
		secondaryColors = determineColors(&color2Flag)
	}

	var output []string

	// Process each part and determine which characters should be colored
	for _, part := range parts {
		if part == "" {
			output = append(output, "")
			continue
		}

		// Create color flags for the current part
		colorSubs := determineSubToColor(part, substring)

		// Generate the ASCII art with the appropriate colors
		result, err := printAsciiArt(part, font, primaryColors, secondaryColors, colorSubs)
		if err != nil {
			fmt.Println(err)
			return
		}
		output = append(output, result...)
	}

	// Print to terminal
	for _, line := range output {
		fmt.Println(line)
	}

	// Write to file if specified
	if outputFile != "" {
		err := writeToFile(output, outputFile)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func writeToFile(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		// Strip ANSI codes for a plain text output
		_, err := file.WriteString(StripAnsiCodes(line) + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
