package piscine

func StrRev(s string) string {
	arr := []rune(s)
	arr1 := []rune(s)
	x := 0
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = arr1[x]
		x++
	}
	return string(arr)
}
