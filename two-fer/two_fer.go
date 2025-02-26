package twofer

// ShareWith returns a string with the message "One for <name>, one for me."
// If no name is given, the result should be "One for you, one for me."
func ShareWith(name string) string {
	if len(name) > 0 {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}