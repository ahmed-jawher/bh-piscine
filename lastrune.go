package piscine

func LastRune(s string) rune {
	f := []rune(s)
	return f[len(s)-1]
}
