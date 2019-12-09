//Package luhn to validate string per the Luhn formula.
package luhn

import (
	"strings"
)

var sum int

//Valid returns if a string is a valid Luhn string
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	//step 1: compute sum
	double := len(s)%2 == 0
	sum = 0
	for _, rV := range s {
		digit := int(rV - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}
		sum += digit
		double = !double
	}

	//step 2: return remainder of sum
	return len(s) > 1 && sum%10 == 0
}
