package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundSign := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '-' && !foundSign {
			sign = -1
			foundSign = true
		} else if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
			foundSign = true
		}
	}
	return result * sign
}
