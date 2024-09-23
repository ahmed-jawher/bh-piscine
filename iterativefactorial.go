package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb == 0 {
		return 1
	}
	if nb > 21 || nb < 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		result = i * result
	}
	return result
}
