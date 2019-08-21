//Package collatzconjecture to solve the Callatz Conjecture problem from exercism
package collatzconjecture

import (
	"errors"
)

func parity(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

func collatz(num int) int {
	par := parity(num)
	if par == "even" {
		return (num / 2)
	}
	return (num * 3) + 1
}

//CollatzConjecture calculates the stoping time of an integer under the collatz operation.
func CollatzConjecture(num int) (int, error) {
	if num <= 0 {
		return -1, errors.New("there was an error")
	}
	count := 0
	for num != 1 {
		num = collatz(num)
		count = count + 1
	}
	return count, nil
}
