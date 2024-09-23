package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	length := len(runes)
	result := make([]rune, 0, length)
	shouldCapitalize := true
	for i := 0; i < length; i++ {
		char := runes[i]
		if isAlphaNumeric(char) {
			if shouldCapitalize {
				char = toUpper(char)
				shouldCapitalize = false
			} else {
				char = toLower(char)
			}
		} else {
			shouldCapitalize = true
		}
		result = append(result, char)
	}
	return string(result)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
