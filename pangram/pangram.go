package pangram

import (
	"strings"
)

//IsPangram check a string is spans the entire ASCII alphabet set
func IsPangram(s string) bool {
	vals := make([]bool, 26)
	for _, rV := range strings.ToLower(s) {
		if rV < 97 || rV > 122 {
			continue
		}
		vals[rV-97] = true
	}
	for _, v := range vals {
		if v != true {
			return false
		}
	}
	return true
}
