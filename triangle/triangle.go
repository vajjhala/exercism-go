// Package triangle determins the kind of triangle
package triangle

// Kind represents the kind of traingle
type Kind string

const (
	// NaT is not a triangle
	NaT Kind = "not"
	// Equ is an equilateral
	Equ Kind = "equ"
	// Iso is an isosceles
	Iso Kind = "iso"
	// Sca is ascalene
	Sca Kind = "sca"
)

// KindFromSides takes the lengths of the three sides of a triangle and returns the Kind of triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k
}
