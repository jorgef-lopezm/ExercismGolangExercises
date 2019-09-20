package hamming

import "errors"

// Distance returns an int with the amount of symbols different on a and b, if lengths are distinct it returns a -1 and a error
func Distance(a, b string) (int, error) {
	if len(a) !=  len(b) {
		return -1, errors.New("The lenght of string a and b are different")
	}

	var counter int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			counter++
		}
	}

	return counter, nil
}
