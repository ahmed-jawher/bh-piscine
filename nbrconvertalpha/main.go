package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Get command-line arguments excluding the program name.

	upperFlag := false

	// Check if the first argument is the --upper flag
	if len(args) > 0 && args[0] == "--upper" {
		upperFlag = true
		args = args[1:] // Remove the flag from arguments
	}

	// Function to convert string to integer manually
	atoi := func(s string) (int, bool) {
		num := 0
		for _, r := range s {
			if r < '0' || r > '9' {
				return 0, false // Return false if non-digit characters are found
			}
			num = num*10 + int(r-'0') // Convert rune to its integer value
		}
		return num, true
	}

	// Loop over each argument
	for _, arg := range args {
		n, valid := atoi(arg) // Try to convert argument to an integer manually.
		if !valid || n < 1 || n > 26 {
			// Print a space for invalid input.
			z01.PrintRune(' ')
		} else {
			letter := rune('a' + n - 1) // Calculate the corresponding letter.
			if upperFlag {
				letter -= 32 // Convert to uppercase by subtracting 32 from the lowercase rune.
			}
			z01.PrintRune(letter) // Print the letter using z01.PrintRune
		}
	}

}
