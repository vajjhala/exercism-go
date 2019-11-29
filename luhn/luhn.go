//Package luhn to validate string per the Luhn formula.
package luhn

import (
	"regexp"
	"strconv"
)

//Valid returns if a string is a valid Luhn string
func Valid(s string) bool {
	//step 0: check if valid Luhn string
	mtch, _ := regexp.MatchString(`^\d(\d|" ")*$`, s)
	if !mtch {
		return false
	}
	re := regexp.MustCompile(`^\d(\d|" ")*$`)
	s = re.ReplaceAllString(" ", "")
	//step 1: compute sum
	runes := []rune(s)
	slen, sum, digit := len(runes), 0, 0
	for i := 1; i < slen; i++ {
		digit, ok := strconv.Atoi(strconv.QuoteRuneToASCII(runes[slen-i]))
		if ok != nil {
			return false
		}
		if i%2 == 0 {
			digit = digit % 9
		}
		sum = sum + digit
	}
	//step 2: check sum is evenly divisible by 10
	return (digit%10 == 0)
}
