package piscine

import (
	"os"
	"github.com/01-edu/z01"
)

func main0() {
	// Check if the program receives exactly one argument.
	if len(os.Args) != 2 {
		z01.PrintRune('\n') // Print a newline if arguments are not exactly one.
		return
	}

	// Get the input string
	input := os.Args[1]
	var result []rune
	wasSpace := false

	// Trim leading spaces and process the string
	for i := 0; i < len(input); i++ {
		char := input[i]

		// Ignore leading spaces
		if char == ' ' || char == '\t' {
			if len(result) > 0 && !wasSpace {
				result = append(result, ' ') // Add a single space if not already added
				wasSpace = true
			}
		} else {
			result = append(result, rune(char)) // Add the non-space character
			wasSpace = false
		}
	}

	// If the result is empty, there's no word to print
	if len(result) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Print the result followed by a newline
	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n') // Add a newline at the end
}
