package isogram

import (
	"strings"
)

//IsIsogram takes (any string) and returns (Isogram or not)
func IsIsogram(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	cnt := make(map[rune]int)
	for _, rV := range s {
		if cnt[rV] != 0 {
			return false
		}
		if (rV != '-') && (rV != ' ') {
			cnt[rV] = 1
		}
	}
	return true
}
