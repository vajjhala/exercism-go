//Package luhn to validate string per the Luhn formula.
package luhn

import (
	//"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Valid returns if a string is a valid Luhn string
func Valid(s string) bool {
	//step 0: check if valid Luhn string
	pattern := `^\d{1}(\d|\s)+$`
	mtch, _ := regexp.MatchString(pattern, s)
	if !mtch {
		return false
	}
	//re := regexp.MustCompile(pattern)
	s = strings.Replace(s, " ", "", -1)
	//step 1: compute sum
	runes := []rune(s)
	slen, sum := len(runes), 0
	//fmt.Printf("\nrunes: %s, len: %d\n", s, slen)
	for i := 1; i <= slen; i++ {
		digit, err := strconv.ParseInt(string(runes[slen-i]), 10, 8)
		if err != nil {
			return false
		}
		if i%2 == 0 {
			digit = (digit * 2)
		}
		if digit > 9 {
			digit = digit - 9
		}

		sum = sum + int(digit)
	}

	//step 2: return remainder of sum
	return (sum%10 == 0)
}
