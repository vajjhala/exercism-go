//Package grains to solve the "grains on a chessboard" problem
package grains

import (
	"errors"
)

//Square returns the number of grains on the square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("number out of chessboard range")
	}
	// start with 1grain
	return 1 << (n - 1), nil
}

//Total returns the maximum number of grains on a chessboard
func Total() uint64 {
	return (1 << 64) - 1
}
