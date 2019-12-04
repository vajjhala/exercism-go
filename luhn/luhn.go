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
	//fmt.Printf("\nrunes: %s, len: %d\n", s, slen)
	for i := 1; i <= slen; i++ {
		digit := int(runes[slen-i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if i%2 == 0 {
			digit = (digit * 2)
		}
		if digit > 9 {
			digit = digit - 9
		}
		//fmt.Printf(", sum = %d; ", sum)
		sum = sum + digit
	}

	//step 2: return remainder of sum
	if (sum == 0) && (slen <= 1) {
		return false
	}
	return (sum%10 == 0)
}
