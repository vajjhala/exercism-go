//Package cryptosquare to encode using crypto square
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

var sq = func(text *string) (c, r, l int) {
	l = len(*text)
	c = int(math.Ceil(math.Sqrt(float64(l))))
	r = c - 1
	if (c * r) >= l {
		return c, r, l
	}
	return c, c, l
}

var normalize = func(r rune) rune {
	if !unicode.In(r, unicode.Letter, unicode.Number) {
		return -1
	}
	return unicode.ToLower(r)
}

//Encode ...
func Encode(message string) string {
	message = strings.Map(normalize, message)
	clm, row, l := sq(&message)
	square := make([]string, clm)
	//add padding
	for n := clm*row - l; n > 0; n-- {
		message += " "
	}
	//increment each row-string
	for i, rV := range message {
		square[i%clm] += string(rV)
	}
	return strings.Join(square, " ")
}
