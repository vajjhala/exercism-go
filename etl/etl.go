//Package etl to transform scores
package etl

import (
	"strings"
)

//Transform (map[int][]string) map[string]int
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for k, arr := range in {
		for _, v := range arr {
			out[strings.ToLower(v)] = k
		}
	}
	return out
}
