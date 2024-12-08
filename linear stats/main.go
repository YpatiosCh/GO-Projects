package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// readData reads numerical data from a file, handling potential errors
// Purpose: Convert file contents into a slice of float64 for analysis
func readData(filename string) []float64 {
	// Open the file, exit if file cannot be accessed
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var data []float64
	scanner := bufio.NewScanner(file)

	// Read each line, convert to float, and store in data slice
	for scanner.Scan() {
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatalf("Invalid data in file: %v", err)
		}
		data = append(data, val)
	}

	return data
}

// calculateMean computes the average of a slice of numbers
// Purpose: Find the central tendency of the dataset
func calculateMean(data []float64) float64 {
	// Sum all values and divide by total count
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

// calculateLinearRegression finds the best-fit line for the dataset
// Purpose: Determine the trend and predictive line of the data
func calculateLinearRegression(x, y []float64) (float64, float64) {
	// Calculate mean values for x and y
	xMean := calculateMean(x)
	yMean := calculateMean(y)

	// Variables to store slope calculation components
	var numerator, denominator float64

	// Compute slope and intercept components
	for i := range x {
		// Calculate deviations from mean for precise line fitting
		numerator += (x[i] - xMean) * (y[i] - yMean)
		denominator += math.Pow(x[i]-xMean, 2)
	}

	// Calculate slope (m) and y-intercept (b)
	m := numerator / denominator
	b := yMean - m*xMean

	return m, b
}

// calculatePearsonCorrelation measures the linear relationship strength
// Purpose: Determine how closely related the x and y variables are
func calculatePearsonCorrelation(x, y []float64) float64 {
	// Calculate mean values for x and y
	xMean := calculateMean(x)
	yMean := calculateMean(y)

	// Variables to store correlation calculation components
	var covariance, xStdDev, yStdDev float64

	// Compute covariance and standard deviations
	for i := range x {
		// Calculate shared variation and individual spread
		covariance += (x[i] - xMean) * (y[i] - yMean)
		xStdDev += math.Pow(x[i]-xMean, 2)
		yStdDev += math.Pow(y[i]-yMean, 2)
	}

	// Calculate correlation coefficient
	return covariance / math.Sqrt(xStdDev*yStdDev)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: program <datafile>")
	}

	data := readData(os.Args[1])

	x := make([]float64, len(data))
	for i := range x {
		x[i] = float64(i)
	}

	m, b := calculateLinearRegression(x, data)

	r := calculatePearsonCorrelation(x, data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
