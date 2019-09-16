//Package scrabble to get the scrabble score of a word
package scrabble

import "strings"

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func mapper() map[string]int {
	var values = make(map[string]int, 26)
	for _, value := range alphabets {
		letter := string(value)
		switch letter {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			values[letter] = 1
		case "d", "g":
			values[letter] = 2
		case "b", "c", "m", "p":
			values[letter] = 3
		case "f", "h", "v", "w", "y":
			values[letter] = 4
		case "k":
			values[letter] = 5
		case "j", "x":
			values[letter] = 8
		case "q", "z":
			values[letter] = 10
		default:
			values[letter] = 0
		}
	}
	return values
}

//Score returns the scrabble score of the word
func Score(word string) int {
	word = strings.ToLower(word)
	values := mapper()
	score := 0
	for i := 0; i < len(word); i++ {
		s, ok := values[string(word[i])]
		if ok != true {
			return 0
		}
		score = score + s
	}
	return score
}
