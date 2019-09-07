// Package triangle determins the kind of triangle
package triangle

import (
	"math"
	"sort"
)

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

// KindFromSides takes the lengths of the three sides of a triangle and returns the Kind of triangle
func KindFromSides(a, b, c float64) Kind {
	var sides = []float64{a, b, c}
	sort.Float64s(sides)
	//conditions to check if it's legal triangle
	if (sides[2] > 0 && sides[2] < math.Inf(1) && sides[0] > 0 && (sides[0]+sides[1]) >= sides[2]) == false {
		return NaT
	}
	if sides[0] == sides[1] && sides[1] == sides[2] {
		return Equ
	}
	if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}
	return Sca
}
