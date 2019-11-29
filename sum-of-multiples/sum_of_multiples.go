package summultiples

//SumMultiples (lmt, divisors) returns int
func SumMultiples(lmt int, divisors ...int) int {
	//fmt.Printf("\nInput: %v, limit: %d\n", divisors, lmt)
	sum := 0
	var uniq = make(map[int]bool, lmt)
	for i := 0; i < len(divisors); i++ {
		num := divisors[i]
		//fmt.Printf("divisor: %d\n", num)
		if num <= 0 {
			continue
		}
		//fmt.Print("\tmultiples: ")
		for j := 1; num*j < lmt; j++ {
			v := num * j
			if uniq[v] {
				continue
			}
			sum = sum + v
			uniq[v] = true
			//fmt.Printf("%d (%d), ", num*j, sum)
		}
		//fmt.Println()
	}
	return sum
}
