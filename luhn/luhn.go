//Package luhn to validate string per the Luhn formula.
package luhn

import (
	"strings"
)

//Valid returns if a string is a valid Luhn string
func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	//step 1: compute sum
	runes := []rune(s)
	slen, sum := len(runes), 0
	double := false
	if slen%2 == 0 {
		double = true
	}
	for _, rV := range runes {
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
	return (slen > 1) && sum%10 == 0
}
