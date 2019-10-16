package encode

import (
	"strconv"
	"strings"
	"unicode"
)

//RunLengthEncode encodes using RLE
func RunLengthEncode(s string) string {
	var b strings.Builder
	for len(s) > 0 {
		cutset := string(s[0])
		lb := len(s)
		s = strings.TrimLeft(s, cutset)
		cutsetLen := lb - len(s)
		if cutsetLen > 1 {
			b.WriteString(strconv.Itoa(cutsetLen))
		}
		b.WriteString(cutset)
	}
	return b.String()
}

//RunLengthDecode decodes using RLE
func RunLengthDecode(s string) string {
	var b strings.Builder
	count := 1
	var prvD rune
	for _, rV := range s {
		stg := string(rV)
		if unicode.IsDigit(rV) {
			if unicode.IsDigit(prvD) {
				stg = strconv.Itoa(count) + stg
			}
			count, _ = strconv.Atoi(stg)
			prvD = rV
			continue
		}
		for k := 0; k < count; k++ {
			b.WriteString(string(rV))
		}
		prvD = rV
		count = 1
	}
	return b.String()
}
