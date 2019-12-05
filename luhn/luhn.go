//Package luhn to validate string per the Luhn formula.
package luhn

import (
	//"fmt"
	//"regexp"
	//"strconv"
	"strings"
)

//Valid returns if a string is a valid Luhn string
func Valid(s string) bool {
	//re := regexp.MustCompile(pattern)
	//fmt.Printf("\n***** string = %s *****", s)
	s = strings.ReplaceAll(s, " ", "")
	//step 1: compute sum
	runes := []rune(s)
	slen, sum := len(runes), 0
	double := false
	//fmt.Printf("\nrunes: %s, len: %d\n", s, slen)
	for i := 1; i <= slen; i++ {
		digit := int(runes[slen-i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit = (digit * 2)
		}
		if digit > 9 {
			digit = digit - 9
		}
		//fmt.Printf(", sum = %d; ", sum)
		sum += digit
		double = !double
	}

	//step 2: return remainder of sum
	return (slen > 1) && sum%10 == 0
}
