package reverse

import "bytes"

// Reverse returns word in reverse
func Reverse(word string) string {
	var output bytes.Buffer

	r := []rune(word)
	for i := len(r) - 1; i >= 0; i-- {
		output.WriteRune(r[i])
	}

	return output.String()
}