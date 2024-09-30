package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	for _, ch := range arg[2:] {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
