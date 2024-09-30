package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:] // Get command-line arguments excluding the program name.

	upperFlag := false

	// Check if the first argument is the --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upperFlag = true
		args = args[1:] // Remove the flag from arguments
	}

	var result strings.Builder

	// Loop over each argument
	for _, arg := range args {
		n, err := strconv.Atoi(arg) // Try to convert argument to an integer.
		if err != nil || n < 1 || n > 26 {
			result.WriteRune(' ') // Write a space for invalid input.
		} else {
			letter := rune('a' + n - 1) // Calculate the corresponding letter.
			if upperFlag {
				letter = unicode.ToUpper(letter) // Convert to uppercase if --upper flag is set.
			}
			result.WriteRune(letter)
		}
	}

	fmt.Println(result.String())
}