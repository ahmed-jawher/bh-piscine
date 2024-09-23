package piscine

func IterativePower(nbr int, power int) int {
	pow := 1
	if power > 0 {
		for i := power; i >= 1; i-- {
			pow *= nbr
		}
	} else if power == 0 {
		pow = 1
	} else if power < 0 {
		pow = 0
	}
	return pow
}
