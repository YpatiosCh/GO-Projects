package asciiTools

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	orange  = "\033[38;5;214m"
	white   = "\033[37m"
)

// Extended color mapping for named colors with their RGB equivalents
var colorMap = map[string][3]int{
	"red":             {255, 0, 0},
	"green":           {0, 255, 0},
	"yellow":          {255, 255, 0},
	"blue":            {0, 0, 255},
	"magenta":         {255, 0, 255},
	"cyan":            {0, 255, 255},
	"orange":          {255, 165, 0},
	"purple":          {128, 0, 128},
	"pink":            {255, 192, 203},
	"brown":           {165, 42, 42},
	"gray":            {128, 128, 128},
	"black":           {0, 0, 0},
	"white":           {255, 255, 255},
	"navy":            {0, 0, 128},
	"teal":            {0, 128, 128},
	"lime":            {0, 255, 0},
	"olive":           {128, 128, 0},
	"gold":            {255, 215, 0},
	"coral":           {255, 127, 80},
	"lightcoral":      {240, 128, 128},
	"darkorange":      {255, 140, 0},
	"lightgreen":      {144, 238, 144},
	"darkgreen":       {0, 100, 0},
	"salmon":          {250, 128, 114},
	"darkred":         {139, 0, 0},
	"crimson":         {220, 20, 60},
	"indigo":          {75, 0, 130},
	"violet":          {238, 130, 238},
	"khaki":           {240, 230, 140},
	"plum":            {221, 160, 221},
	"slategray":       {112, 128, 144},
	"lightslategray":  {119, 136, 153},
	"lightsalmon":     {255, 160, 122},
	"lightseagreen":   {32, 178, 170},
	"mediumseagreen":  {60, 179, 113},
	"mediumblue":      {0, 0, 205},
	"mediumvioletred": {199, 21, 133},
	"peru":            {205, 133, 63},
	"darkslategray":   {47, 79, 79},
	"chocolate":       {210, 105, 30},
	"mediumturquoise": {72, 209, 204},
	"lightblue":       {173, 216, 230},
	"darkviolet":      {148, 0, 211},
	"lawngreen":       {124, 252, 0},
	"forestgreen":     {34, 139, 34},
	"seashell":        {255, 228, 225},
	"antiquewhite":    {250, 235, 215},
	"lightpink":       {255, 182, 193},
	"darkkhaki":       {189, 183, 107},
	"goldenrod":       {218, 165, 32},
	"mediumorchid":    {186, 85, 211},
	"lightgray":       {211, 211, 211},
	"dimgray":         {105, 105, 105},
}

// determineColors gets the colors needed from the flag --color=
// It returns a slice of color strings
// determineColor gets the color needed from the flag --color=
func determineColors(colorFlag *string) []string {
	// Split the input colors by comma
	colorNames := strings.Split(*colorFlag, ",")
	colors := make([]string, len(colorNames))

	for i, color := range colorNames {
		color = strings.TrimSpace(color)
		switch {
		case strings.HasPrefix(color, "#"): // Hex format
			colors[i] = hexToAnsi(color)
		case strings.HasPrefix(color, "rgb("): // RGB format
			colors[i] = rgbToAnsi(color)
		case strings.HasPrefix(color, "hsl("): // HSL format
			colors[i] = hslToAnsi(color)
		default: // Named colors
			colors[i] = namedColorToAnsi(color)
		}
	}

	return colors
}

// Helper function to convert named colors to ANSI codes
func namedColorToAnsi(color string) string {
	if rgb, exists := colorMap[color]; exists {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb[0], rgb[1], rgb[2])
	}
	return white // Fallback to white for invalid named colors
}

// Function to convert hex color code to ANSI color code
func hexToAnsi(hex string) string {
	if len(hex) != 7 || hex[0] != '#' {
		return white // Fallback to white for invalid hex
	}

	r, _ := strconv.ParseInt(hex[1:3], 16, 0)
	g, _ := strconv.ParseInt(hex[3:5], 16, 0)
	b, _ := strconv.ParseInt(hex[5:7], 16, 0)

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// Function to convert RGB color format to ANSI color code
func rgbToAnsi(rgb string) string {
	re := regexp.MustCompile(`rgb\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)`)
	matches := re.FindStringSubmatch(rgb)
	if len(matches) == 4 {
		r, _ := strconv.Atoi(matches[1])
		g, _ := strconv.Atoi(matches[2])
		b, _ := strconv.Atoi(matches[3])
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	}
	return white // Fallback to white for invalid RGB
}

// Function to convert HSL color format to ANSI color code
func hslToAnsi(hsl string) string {
	re := regexp.MustCompile(`hsl\(\s*(\d+)\s*,\s*(\d+)%\s*,\s*(\d+)%\s*\)`)
	matches := re.FindStringSubmatch(hsl)
	if len(matches) == 4 {
		h, _ := strconv.Atoi(matches[1])
		s, _ := strconv.Atoi(matches[2])
		l, _ := strconv.Atoi(matches[3])

		// Convert HSL to RGB
		r, g, b := hslToRgb(h, s, l)
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	}
	return white // Fallback to white for invalid HSL
}

// Helper function to convert HSL to RGB
func hslToRgb(h, s, l int) (int, int, int) {
	var r, g, b float64
	sF := float64(s) / 100 // Convert to float and normalize
	lF := float64(l) / 100 // Convert to float and normalize

	if sF == 0 {
		// Achromatic
		r, g, b = lF, lF, lF
	} else {
		var q float64
		if lF < 0.5 {
			q = lF * (1 + sF)
		} else {
			q = lF + sF - lF*sF
		}
		p := 2*lF - q

		r = hueToRgb(p, q, (float64(h)/360)+1.0/3.0)
		g = hueToRgb(p, q, float64(h)/360)
		b = hueToRgb(p, q, (float64(h)/360)-1.0/3.0)
	}

	// Convert back to int and return
	return int(r * 255), int(g * 255), int(b * 255)
}

// Helper function for HSL to RGB conversion
func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1/6 {
		return p + (q-p)*6*t
	}
	if t < 1/2 {
		return q
	}
	if t < 2/3 {
		return p + (q-p)*(2/3-t)*6
	}
	return p
}

// Function to determine which characters should be colored
func determineSubToColor(input string, substring string) []bool {
	// Handle edge cases
	if len(substring) == 0 || len(substring) > len(input) {
		return make([]bool, len(input)) // Return all false if the substring is invalid
	}

	colorFlags := make([]bool, len(input))

	// Find occurrences of the substring
	for i := 0; i <= len(input)-len(substring); i++ {
		if input[i:i+len(substring)] == substring {
			// Mark the indices for coloring
			for j := 0; j < len(substring); j++ {
				colorFlags[i+j] = true
			}
		}
	}

	return colorFlags
}

// Function to strip ANSI color codes from strings
func StripAnsiCodes(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}

// Validate the color flag to ensure it uses "=" syntax
func ValidateColorFlag() error {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--color") && !strings.Contains(arg, "=") {
			return fmt.Errorf("invalid syntax for --color flag. please use '--color=colorname' format")
		}
	}
	return nil
}
