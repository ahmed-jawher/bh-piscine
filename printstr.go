package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	arr := []rune(s)
	for i := 0; i <= len(arr)-1; i++ {
		z01.PrintRune(arr[i])
	}
}
