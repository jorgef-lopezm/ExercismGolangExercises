package bob

import (
	"strings"
	"unicode"
)

func isAYell (input string) bool {
	// Making sure it has at least one letter, if no letter found IndexFunc returns -1
	if !(strings.IndexFunc(input, unicode.IsLetter) >= 0) {
		return false
	}

	for _, symbol := range input {
		if symbol > 'a' && symbol < 'z' {
			return false
		}
	}
	return true
}

func isAQuestion (input string) bool {
	return input[len(input) - 1] == '?'
}

func isYellingAQuestion (input string) bool {
	if isAYell(input) && isAQuestion(input) {
		return true
	}
	return false
}

// Hey returns a string depending on the parameter.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	if isYellingAQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isAYell(remark) {
		return "Whoa, chill out!"
	}

	if isAQuestion(remark) {
		return "Sure."
	}

	return "Whatever."
}
