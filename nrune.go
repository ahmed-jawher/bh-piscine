package piscine

func NRune(s string, n int) rune {
	w := []rune(s)
	if n <= 0 || n > len(w) {
		return 0
	} else {
		return w[n-1]
	}
}
