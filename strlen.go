package piscine

func StrLen(s string) int {
	arr := []rune(s)
	y := 0
	for i := 0; i <= len(arr)-1; i++ {
		y++
	}
	return y
}
