package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		n = -n
	}
	numDigits := countDigits(n)
	digits := make([]int, numDigits)
	for i := numDigits - 1; i >= 0; i-- {
		digits[i] = n % 10
		n /= 10
	}
	for i := 0; i < numDigits-1; i++ {
		for j := 0; j < numDigits-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
	for _, digit := range digits {
		z01.PrintRune(rune(digit) + '0')
	}
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
