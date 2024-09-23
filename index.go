package piscine

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:len(toFind)+i] == toFind {
			return i
		}
	}
	return -1
}
