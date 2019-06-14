package diffsquares

//SquareOfSum returns the square of the sum of natural numbers up to the input
func SquareOfSum(num int) int {
	n := ((num * (num + 1)) / 2)
	return n * n
}

//SumOfSquares returns the sum of the squares of natural numbers up to the input
func SumOfSquares(num int) int {
	return (num * (num + 1) * ((2 * num) + 1) / 6)
}

//Difference returns the differences between Square of Sum and Sum of Squares
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
