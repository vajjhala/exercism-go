package scrabble

import "strings"

type one string
type two string
type three string
type four string
type five string
type eight string
type ten string

const (
	a, e, i, o, u, l, n, r, s, t one   = "a", "e", "i", "o", "u", "l", "n", "r", "s", "t"
	d, g                         two   = "d", "g"
	b, c, m, p                   three = "b", "c", "m", "p"
	f, h, v, w, y                four  = "f", "h", "v", "w", "y"
	k                            five  = "k"
	j, x                         eight = "j", "k"
	q, z                         ten   = "q", "z"
)

func alphabetScores()

func Score(word string) int {
	scoreMap := make(map[string]int, 26)

}
