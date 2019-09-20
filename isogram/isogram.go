package isogram

import "strings"

// IsIsogram returns a bool checking if word is an isogram or not
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for _, symbol := range word {
		if symbol == ' ' || symbol == '-' {
			continue
		}

		if (strings.Count(word, string(symbol)) > 1) {
			return false
		}
	}

	return true
}