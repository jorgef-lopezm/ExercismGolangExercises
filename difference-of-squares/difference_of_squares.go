package diffsquares

func SquareOfSum(number int) (result int) {
	result = number * (number + 1) / 2
	return result * result
}

func SumOfSquares(number int) int {
	return (number * (number + 1) * (2 * number + 1)) / 6

}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}