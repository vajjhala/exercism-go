//Package clock to add, subtract and display time
package clock

import (
	"fmt"
)

//Clock abstracts time as a 24hr display clock
type Clock struct {
	minute int
}

//clock converts an any integer to a clock value
func (c Clock) clock() Clock {
	// clock value range: [0,1439]
	// 1440 are the number of minutes in a day
	c.minute %= 1440
	if c.minute < 0 {
		c.minute += 1440
	}
	return c
}

//New takes in hours and minutes to create a clock value
func New(hour, minute int) Clock {
	return Clock{minute + hour*60}.clock()
}

//String formats the display for a clock value
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}

//Add to add minutes to a clock value
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.clock()
}

//Subtract minutes from a clock value
func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	return c.clock()
}
