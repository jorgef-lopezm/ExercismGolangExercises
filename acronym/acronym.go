package acronym

import "strings"

type SplittedPhrase struct {
	words []string
}

// Returns a SplittedPhrase struct, it removes all unwanted symbols from s and splits s.
func newSplittedPhrase(s string) SplittedPhrase {
	s = strings.Replace(s, "_", "", -1)
	s = strings.Replace(s, "-", " ", -1)
	
	return SplittedPhrase { strings.Fields(s) }
}

// Returns a string with the initial symbol of each string on slice
func getInitials(words []string) string {
	var acronym string = ""

	for _, w := range words {
		acronym += string(w[0])
	}

	return strings.ToUpper(acronym)
}

// Abbreviate returns a string with acronym for s
func Abbreviate(s string) string {
	splittedPhrase := newSplittedPhrase(s)

	return getInitials(splittedPhrase.words)
}
