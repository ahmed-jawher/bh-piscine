package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	insertValue := ""
	orderFlag := false
	stringArg := ""

	for i, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertValue = arg[9:]
		} else if arg == "-i" {
			if i+1 < len(args) {
				insertValue = args[i+1]
			}
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
		} else {
			stringArg = arg
		}
	}

	if insertValue != "" {
		stringArg += insertValue
	}

	if orderFlag {
		runes := []rune(stringArg)
		sortRunes(runes)
		stringArg = string(runes)
	}

	for _, r := range stringArg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	helpText := `--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.
`
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
