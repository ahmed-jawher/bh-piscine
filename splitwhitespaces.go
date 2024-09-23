package piscine

func SplitWhiteSpaces(s string) []string {
	var str []string
	var temp string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			temp += string(s[i])
		} else if temp != "" {
			str = append(str, temp)
			temp = ""
		}
	}
	if temp != "" {
		str = append(str, temp)
	}
	return str
}
