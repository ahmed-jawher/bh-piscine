package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	// Special case for zero
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Convert the number to a slice of runes
	digits := []rune{}

	for n > 0 {
		digit := n % 10
		digits = append([]rune{rune('0' + digit)}, digits...) // prepend the digit to the slice
		n /= 10
	}

	// Use for range to print each rune in the slice
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}
