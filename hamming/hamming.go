// Package hamming to solve the Hamming Distance between two DNA strangs
package hamming

import "errors"

func hammingDist(a, b string) int {
	count := 0

	return count
}

//Distance calculates the hamming distance between two DNA strangs and returns an integer and the errors.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The lengths of the two DNA strands must be equal")
	}
	if a == b {
		return 0, nil
	}

	count := hammingDist(a, b)
	return count, nil
}
