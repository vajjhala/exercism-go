//Package diffsquares to calculate the difference between the sum of squares and
//square of sums
package diffsquares

//SquareOfSum calculates (1 + 2 + 3 + ... N)^2
func SquareOfSum(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}

//SumOfSquares calculates 1^2 + 2^2 + 3^2 + ... N^2
func SumOfSquares(n int) int {
	return n * (2*n + 1) * (n + 1) / 6
}

//Difference calculates (1 + 2 + 3 + ... N)^2  -   (1^2 + 2^2 + 3^2 + ... N^2)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
