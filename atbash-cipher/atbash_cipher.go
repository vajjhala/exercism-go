//Package atbash ...
package atbash

import (
	"strings"
	"unicode"
)

func atbash(x rune) rune {
	switch {
	case unicode.IsNumber(x):
		return x
	case unicode.IsLetter(x):
		return 25 - unicode.ToLower(x) + 97 + 97
	default:
		return -1
	}
}

//Atbash ...
func Atbash(plain string) (cipher string) {
	for i, rV := range strings.Map(atbash, plain) {
		if i != 0 && i%5 == 0 {
			cipher += " "
		}
		cipher += string(rV)
	}
	return cipher
}
