package luhn

import "fmt"

func Valid(creditCardNumber string) bool {
	var result, counter, temp int = 0, 1, 0
	

	for i := len(creditCardNumber) - 1; i >= 0; i-- {
		symbol := creditCardNumber[i]

		if (symbol == ' ') {
			continue
		}

		if (symbol >= '0' && symbol <= '9') {
			if (counter % 2 == 0) {
				temp = (2 * int(symbol - '0'))

				if (temp > 9) {
					temp -= 9
				}
			} else {
				temp = int(symbol - '0')
			}
			result += temp
			counter++
		} else {
			return false
		}
	}
	fmt.Println(counter, result)
	return (counter > 2) && (result % 10 == 0)
}