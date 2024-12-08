package asciiTools

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// HandleArgs processes and validates command-line arguments, returning parameters or an error if invalid
func HandleArgs() (string, string, string, string, string, string, error) {
	// Define flags
	colorFlag := flag.String("color", "white", "Specify the output color (e.g., red, green, yellow)")
	color2Flag := flag.String("color2", "", "Specify a second output color (optional)")
	outputFlag := flag.String("output", "", "Output file (e.g., output.txt)")

	// Parse flags
	flag.Parse()

	// Validate the color flag format
	if err := ValidateColorFlag(); err != nil {
		return "", "", "", "", "", "", err
	}

	// Validate output flag if provided
	if *outputFlag != "" && !strings.HasSuffix(*outputFlag, ".txt") {
		return "", "", "", "", "", "", fmt.Errorf("output file must have .txt extension")
	}

	// Check if we have enough arguments
	if flag.NArg() < 1 {
		PrintUsage()
		return "", "", "", "", "", "", errors.New("not enough arguments")
	}

	// Default values
	substring := flag.Arg(0)       // First argument is the substring to be highlighted
	fullInput := substring         // Default fullInput is the same as substring if only one argument is provided
	font := "./fonts/standard.txt" // Default font path

	// Handle argument variations
	switch flag.NArg() {
	case 1:
		// Only the substring provided, use default font and substring as fullInput
		fullInput = substring
	case 2:
		if IsSubstring(flag.Arg(1), flag.Arg(0)) {
			// Two arguments, with second as full input
			substring = flag.Arg(0)
			fullInput = flag.Arg(1)
		} else if !IsSubstring(flag.Arg(1), flag.Arg(0)) && !IsFont(flag.Arg(1)) { // if first arg is not a substring and the second arg is not a fontName
			substring = flag.Arg(0)
			fullInput = flag.Arg(1)
		} else {
			// Second argument specifies a font
			fullInput = flag.Arg(0)
			font = "./fonts/" + flag.Arg(1) + ".txt"
		}
	case 3:
		// Three arguments: substring, input, and font name
		substring = flag.Arg(0)
		fullInput = flag.Arg(1)
		font = "./fonts/" + flag.Arg(2) + ".txt"
	default:
		PrintUsage()
		return "", "", "", "", "", "", errors.New("invalid number of arguments")
	}

	if len(fullInput) == 0 {
		return "", "", "", "", "", "", errors.New("input string cannot be empty")
	}

	return substring, fullInput, *colorFlag, *color2Flag, font, *outputFlag, nil // Updated to return 7 values
}

// PrintUsage prints the usage instructions for the program
func PrintUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("\nEX: go run . --color=<color> <substring to be colored> \"something\"")
}

// // IsFont returns true if the string we check is a fontName
func IsFont(s string) bool {
	return s == "standard" || s == "shadow" || s == "thinkertoy"
}

// IsSubstring function determines whether a string is a substring of another string
func IsSubstring(str, subStr string) bool {
	return strings.Contains(str, subStr)
}
