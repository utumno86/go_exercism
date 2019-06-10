package diffsquares

import "math"

//SquareOfSum returns the square of the sum of natural numbers up to the input
func SquareOfSum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum = sum + i
	}
	return int(math.Pow(float64(sum), 2))
}

//SumOfSquares returns the sum of the squares of natural numbers up to the input
func SumOfSquares(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum = sum + int(math.Pow(float64(i), 2))
	}
	return sum
}

//Difference returns the differences between Square of Sum and Sum of Squares
func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
