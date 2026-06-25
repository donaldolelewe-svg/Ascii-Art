package main

import (
	"fmt"
	"os"
	"strings"
)

// Constant banner file
const bannerFile = "standard.txt"

// Function to load banner file
func loadBanner(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	banner := strings.Split(string(data), "\n")
	return banner, nil
}

// Function to render ASCII art
func render(input string, banner []string) string {
	var result strings.Builder

	// Split input by actual newline (\n from terminal)
	input = strings.ReplaceAll(input, "\\n", "\\t\\n\\t")
	lines := strings.Split(input, "\\t")

	for i, line := range lines {

		// Handle empty line
		if line == "" {
			continue
		}
		if line == "\\n" {
			result.WriteString("\n")
			continue
		}

		// Each character has 8 rows
		for row := 0; row < 8; row++ {

			for _, char := range line {

				// Skip invalid characters
				if char < 32 || char > 126 {
					continue
				}
				// Calculate position in banner
				index := (int(char)-32)*9 + 1

				// Append correct row
				result.WriteString(banner[index+row])
			}
			if i+1 < len(lines) && row > 6 {
				continue
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}

func main() {

	// Check argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"text\"")
		return
	}

	input := os.Args[1]

	// Load banner
	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Render ASCII art
	output := render(input, banner)

	// Print result
	fmt.Print(output)
}
