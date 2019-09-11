//Package proverb generates amazing proverbs.
package proverb

import "fmt"

// Proverb constructs a meaningful proverb.
func Proverb(rhyme []string) []string {
	x := len(rhyme)
	if x < 1 {
		return []string{}
	}
	if x == 1 {
		v := fmt.Printf("And all for the want of a %s", rhyme[0])
		return []string{v}
	}
	proverb = make([]string, len(rhyme), len(rhyme))
	for i := 0; i < len(rhyme); i++ {

	}

}
