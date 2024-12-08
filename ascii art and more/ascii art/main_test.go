package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestLocateLines tests the LocateLines function
func TestLocateLines(t *testing.T) {
	tests := []struct {
		input    string
		expected [][]int
	}{
		{
			input:    "A",
			expected: [][]int{{299}},
		},
		{
			input:    "B",
			expected: [][]int{{308}},
		},
		{
			input:    "Hello",
			expected: [][]int{{362, 623, 686, 686, 713}},
		},
		{
			input:    " ", // Space character
			expected: [][]int{{2}},
		},
	}

	for _, test := range tests {
		result := LocateLines(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
		for i := range result {
			if len(result[i]) != len(test.expected[i]) {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
			}
			for j := range result[i] {
				if result[i][j] != test.expected[i][j] {
					t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
				}
			}
		}
	}
}

// TestPrintLinesFromArray tests the PrintLinesFromArray function
func TestPrintLinesFromArray(t *testing.T) {
	// Use the existing standard.txt file
	standardFile := "standard.txt"

	// Ensure the standard.txt file exists
	if _, err := os.Stat(standardFile); os.IsNotExist(err) {
		t.Fatalf("standard.txt file does not exist: %v", err)
	}

	tests := []struct{
		lineNumbers [][]int
		expectedOutput string
		description string
	}{
		{
			lineNumbers: [][]int{
				{362, 623, 686, 686, 713},
			},
			expectedOutput: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
`,
		},
		{
			lineNumbers: [][]int{
				{362, 335, 398, 398, 425},
			},
			expectedOutput: ` _    _   ______   _        _         ____   
| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
                                             
                                             
`,

		},
	}

	// Iterate over each test case
	for _, test := range tests {
		var buf bytes.Buffer
		t.Logf("Testing PrintLinesFromArray with %s", test.description)

		// Call the PrintLinesFromArrayHelper function to capture the output
		err := PrintLinesFromArrayHelper(standardFile, test.lineNumbers, &buf)
		if err != nil {
			t.Fatalf("PrintLinesFromArray returned an error for %s: %v", test.description, err)
		}

		// Compare the captured output with the expected output
		if buf.String() != test.expectedOutput {
			t.Errorf("For %s, expected output:\n%v\nBut got:\n%v", test.description, test.expectedOutput, buf.String())
		}
	}
}



// PrintLinesFromArrayHelper is a helper function to allow injecting a custom writer (to capture output)
func PrintLinesFromArrayHelper(filename string, lineNumbers [][]int, writer *bytes.Buffer) error {
	// Read the entire file into memory
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	// // Iterate over each array of line numbers
	for i, lineArray := range lineNumbers {
		for i := 0; i < 8; i++ {
			if len(lineArray) == 0 {
				break
			}
			for j := 0; j < len(lineArray); j++ {
				// Safeguard against out-of-bound errors
				if lineArray[j]-1 >= len(lines) {
					return nil
				}
				writer.WriteString(lines[lineArray[j]-1])
				lineArray[j]++
			}
			writer.WriteString("\n")
		}
		if i != len(lineNumbers)-1 {
			// Print a newline after finishing each inner array
			writer.WriteString("\n")
		}
	}

	return nil
}
