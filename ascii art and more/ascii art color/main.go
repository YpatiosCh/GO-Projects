package main

import (
	"ascii-art-color/asciiTools"
	"os"
)

func main() {
	substring, fullInput, colorFlag, color2Flag, font, outputFile, err := asciiTools.HandleArgs() // Add color2Flag variable
	if err != nil {
		asciiTools.PrintUsage()
		os.Exit(1)
	}

	// Pass both colorFlag and color2Flag to PrintFullAscii
	asciiTools.PrintFullAscii(substring, fullInput, colorFlag, color2Flag, font, outputFile)
}
