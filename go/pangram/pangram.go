package pangram

import "unicode"

func IsPangram(input string) bool {
	const lenAlphabets = 26

	m := make(map[rune]bool)
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		m[unicode.ToLower(r)] = true
	}

	if len(m) == lenAlphabets {
		return true
	} else {
		return false
	}
}
