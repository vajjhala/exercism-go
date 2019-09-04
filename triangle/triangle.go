// Package triangle
package triangle

type Kind


const (
    // Pick values for the following identifiers used by the test program.
    NaT // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides takes the lengths of the three sides of a triangle and returns the Kind of triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k
}
