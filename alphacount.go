package piscine

func AlphaCount(s string) int {
	count := 0
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' {
			count++
		}
	}
	return count
}
