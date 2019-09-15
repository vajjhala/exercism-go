//Package proverb generates amazing proverbs.
package proverb

import "fmt"

// Proverb constructs a meaningful proverb.
func Proverb(rhyme []string) []string {
	x := len(rhyme)
	if x < 1 {
		return []string{}
	}
	lastLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	if x == 1 {
		return []string{lastLine}
	}
	proverb := []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	proverb = append(proverb, lastLine)
	return proverb
}
