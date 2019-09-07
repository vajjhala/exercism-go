// Package triangle determins the kind of triangle
package triangle

import "sort"

//Kind is an alias for a string
type Kind string

const (
	// NaT is not a triangle
	NaT = "not"
	// Equ is an equilateral
	Equ = "equ"
	// Iso is an isosceles
	Iso = "iso"
	// Sca is ascalene
	Sca = "sca"
)

func isTriangle(tri []float64) bool {
	if len(tri) != 3 {
		return false
	}
	if tri[0] <= 0 || tri[1] <= 0 || tri[2] <= 0 {
		return false
	}
	sumOfsides := tri[0] + tri[1] + tri[2]
	if sumOfsides <= 0 {
		return false
	}
	for i := 0; i < 3; i++ {
		if tri[i] > tri[i-1]+tri[i-2] {
			return false
		}
	}
	return true
}

// KindFromSides takes the lengths of the three sides of a triangle and returns the Kind of triangle
func KindFromSides(a, b, c float64) Kind {
	var sides = []float64{a, b, c}
	if isTriangle(sides) != true {
		return NaT
	}
	sort.Float64s(sides)
	if sides[1] == sides[2] {
		if sides[0] == sides[1] {
			return Equ
		}
		return Iso
	}
	return Sca
}
