package isogram

import "unicode"

func IsIsogram(word string) bool {
	found := map[rune]bool{}
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}

		r = unicode.ToLower(r)
		if found[r] {
			return false
		}

		found[r] = true
	}
	return true
}
