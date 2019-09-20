package raindrops

import "strconv"

// Convert returns a string with an int converted depending on number
func Convert(number int) (str string) {
	if number % 3 == 0 {
		str += "Pling"
	}

	if number % 5 == 0 {
		str += "Plang"
	}

	if number % 7 == 0 {
		str += "Plong"
	}

	if str == "" {
		return strconv.Itoa(number)
	}
	return str
}