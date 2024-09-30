package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// If no arguments or --help or -h is provided, print the help message.
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insertValue := ""
	orderFlag := false
	stringArg := ""

	// Parse the arguments
	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertValue = arg[9:] // Get the value after --insert=
		} else if arg == "-i" || arg == "--insert" {
			// Expect the next argument to be the value for --insert or -i
			insertValue = args[len(args)-1] // Handles in case there's an insert value
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			stringArg = arg
		}
	}

	// Insert the string if --insert flag is used
	if insertValue != "" {
		stringArg += insertValue
	}

	// Sort the string if --order flag is used
	if orderFlag {
		runes := []rune(stringArg)
		sortRunes(runes)
		stringArg = string(runes)
	}

	// Print the resulting string
	for _, r := range stringArg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	helpText := "--insert\n  -i\n         This flag inserts the string into the string passed as argument.\n--order\n  -o\n         This flag will behave like a boolean, if it is called it will order the argument.\n"
	for _, r := range helpText {
		z01.PrintRune(r)
	}
}

func sortRunes(runes []rune) {
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
}
