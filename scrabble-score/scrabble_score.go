//Package scrabble to get the scrabble score of a word
package scrabble

import "strings"

var values = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

//Score returns the scrabble score of the word
func Score(word string) int {
	score := 0
	for _, runeValue := range strings.ToLower(word) {
		s, ok := values[runeValue]
		if ok != true {
			return 0
		}
		score = score + s
	}
	return score
}
