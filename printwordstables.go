package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, v := range table {
		for _, c := range v {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
