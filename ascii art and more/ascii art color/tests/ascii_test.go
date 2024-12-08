package tests

import (
	"flag"
	"os"
	"strings"
	"testing"

	"ascii-art-color/asciiTools" // Change this to the actual import path of your project
)

// ResetFlags clears existing flags to avoid redefinition
func ResetFlags() {
	// Reset flags to ensure no flags are redefined
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
}

func TestHandleArgs(t *testing.T) {
	tests := []struct {
		args          []string
		expected      []string
		expectedError bool
	}{
		{
			args:          []string{"--color=red", "highlight", "some string", "standard"},
			expected:      []string{"highlight", "some string", "red", "", "./fonts/standard.txt", ""},
			expectedError: false,
		},
		{
			args:          []string{"--color=blue", "highlight"},
			expected:      []string{"highlight", "highlight", "blue", "", "./fonts/standard.txt", ""},
			expectedError: false,
		},
		{
			args:          []string{"--color=green", "highlight", "some string", "shadow"},
			expected:      []string{"highlight", "some string", "green", "", "./fonts/shadow.txt", ""},
			expectedError: false,
		},
		{
			args:          []string{"--color=green", "highlight", "some string", "invalid-font"},
			expectedError: true,
		},
		{
			args:          []string{"--color=green"},
			expectedError: true,
		},
		{
			args:          []string{"--color=green", "--output=output.txt", "highlight", "some string"},
			expected:      []string{"highlight", "some string", "green", "", "./fonts/standard.txt", "output.txt"},
			expectedError: false,
		},
		{
			args:          []string{"--color=green", "--output=output"},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			ResetFlags() // Reset the flags at the start of each test

			// Temporarily set the command-line arguments
			os.Args = append([]string{"cmd"}, tt.args...) // "cmd" is a placeholder for the command name

			substring, fullInput, color, color2, font, output, err := asciiTools.HandleArgs()

			// Check for expected error
			if (err != nil) != tt.expectedError {
				t.Fatalf("expected error: %v, got: %v", tt.expectedError, err)
			}

			if !tt.expectedError {
				// Check expected values if no error is expected
				if substring != tt.expected[0] {
					t.Errorf("expected substring: %s, got: %s", tt.expected[0], substring)
				}
				if fullInput != tt.expected[1] {
					t.Errorf("expected fullInput: %s, got: %s", tt.expected[1], fullInput)
				}
				if color != tt.expected[2] {
					t.Errorf("expected color: %s, got: %s", tt.expected[2], color)
				}
				if color2 != tt.expected[3] {
					t.Errorf("expected color2: %s, got: %s", tt.expected[3], color2)
				}
				if font != tt.expected[4] {
					t.Errorf("expected font: %s, got: %s", tt.expected[4], font)
				}
				if output != tt.expected[5] {
					t.Errorf("expected output: %s, got: %s", tt.expected[5], output)
				}
			}
		})
	}
}

// TestValidateColorFlag tests the ValidateColorFlag function
func TestValidateColorFlag(t *testing.T) {
	// Test valid color flag
	os.Args = []string{"cmd", "--color=red"}
	if err := asciiTools.ValidateColorFlag(); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}

	// Test invalid color flag
	os.Args = []string{"cmd", "--color", "red"}
	if err := asciiTools.ValidateColorFlag(); err == nil {
		t.Errorf("expected an error, got none")
	}

	// Test with no color flag
	os.Args = []string{"cmd"}
	if err := asciiTools.ValidateColorFlag(); err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
}
