package main

import (
	"fmt"
	"os"
)

// Keep track of our current prediction level
var currentLevel float64
var isFirstNumber bool = true

func calculatePredictionLevel(newValue float64) float64 {
	// For the first number, just use it directly
	if isFirstNumber {
		isFirstNumber = false
		currentLevel = newValue
		return currentLevel
	}

	// For all other numbers:
	// Move 15% of the way from our current level towards the new value
	// Example: if currentLevel = 100, newValue = 200
	// difference = 200 - 100 = 100
	// adjustment = 100 * 0.15 = 15
	// new currentLevel = 100 + 15 = 115
	difference := newValue - currentLevel
	adjustment := difference * 0.15
	currentLevel = currentLevel + adjustment

	return currentLevel
}

func main() {
	// Read numbers one by one
	for {
		// Read a number
		var newNumber float64
		_, err := fmt.Fscan(os.Stdin, &newNumber)
		if err != nil {
			break
		}

		// Skip prediction for the first number (as per problem requirements)
		if !isFirstNumber {
			// Calculate our prediction level
			predictedLevel := calculatePredictionLevel(newNumber)

			// Create range of ±6 around our prediction
			// Example: if predictedLevel = 100, range will be 94 to 106
			lowerBound := predictedLevel - 6
			upperBound := predictedLevel + 6

			// Output the prediction range
			fmt.Printf("%.2f %.2f\n", lowerBound, upperBound)
		} else {
			// Just calculate the level for the first number
			calculatePredictionLevel(newNumber)
		}
	}
}
