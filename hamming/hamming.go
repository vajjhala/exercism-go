// Package hamming to solve the Hamming Distance between two DNA strangs
package hamming

import "errors"

func hammingDist(a, b string) int {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count = count + 1
		}
	}
	return count
}

//Distance calculates the hamming distance between two DNA strangs and returns an integer and the errors.
func Distance(a, b string) (int, error) {
	if len(a) == len(b) {
		return hammingDist(a, b), nil
	}
	return 0, errors.New("Error")
}
