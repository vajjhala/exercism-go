package isogram

import (
	"strings"
)

//IsIsogram takes (any string) and returns (Isogram or not)
func IsIsogram(s string) bool {
	cnt := make(map[rune]bool)
	for _, rV := range strings.ToLower(s) {
		if cnt[rV] {
			return false
		}
		if rV == '-' || rV == ' ' {
			continue
		}
		cnt[rV] = true
	}
	return true
}
