package summultiples

import "fmt"

//SumMultiples (lmt, divisors) returns int
func SumMultiples(lmt int, divisors ...int) int {
	sum := 0
	for i := 0; i < len(divisors); i++ {
		if divisors[i] <= 0 {
			continue
		}
		j := 0
		var val int
		val = divisors[i] * j
		for val < lmt {
			fmt.Printf("divisor * j < lmt |  %d (%d * %d) < %d\n",
				val, divisors[i], j, lmt)
			sum = sum + val
			val = divisors[i] * j
			j++
		}
		fmt.Println()
	}
	return sum
}
