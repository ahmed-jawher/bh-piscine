package piscine

func MakeRange(min int, max int) []int {
	if min >= max {
		return nil
	}
	length := max - min
	arr := make([]int, max-min)
	for i := 0; i < length; i++ {
		arr[i] = min + i
	}
	return arr
}
