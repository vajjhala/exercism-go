//Package scrabble to get the scrabble score of a word
package scrabble

import "strings"

const alphabets = "abcdefghijklmnopqrstuvwxyz"

var values = make(map[rune]int, 26)

func init() {
	for _, runeValue := range alphabets {
		letter := string(runeValue)
		switch letter {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			values[runeValue] = 1
		case "d", "g":
			values[runeValue] = 2
		case "b", "c", "m", "p":
			values[runeValue] = 3
		case "f", "h", "v", "w", "y":
			values[runeValue] = 4
		case "k":
			values[runeValue] = 5
		case "j", "x":
			values[runeValue] = 8
		case "q", "z":
			values[runeValue] = 10
		default:
			values[runeValue] = 0
		}
	}
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
