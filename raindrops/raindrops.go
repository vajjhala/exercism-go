package raindrops

import "strconv"

/*
Convert takes an integrer and returns:
    'Pling' or 'Plang' or 'Plong' ?
*/
func Convert(a int) string {
	s, flag := "", false
	if a%3 == 0 {
		s = s + "Pling"
		flag = true
	}
	if a%5 == 0 {
		s = s + "Plang"
		flag = true
	}
	if a%7 == 0 {
		s = s + "Plong"
		flag = true
	}
	if flag == false {
		s = strconv.Itoa(a)
	}
	return s
}
