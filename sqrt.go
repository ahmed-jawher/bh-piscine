package piscine

func Sqrt(nbr int) int {
	if nbr < 0{
	return 0
	}

	for i:=nbr ;i>=0;i--{
		if i*i==nbr{
			return i
		}
	}
	return 0
}
