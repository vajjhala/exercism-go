//Package clock ..
package clock

import (
	// "fmt"
	"strconv"
)

//Clock ...
type Clock struct {
	hour, minute int
}

//clock ..
func (c Clock) clock() Clock {
	x, y, z := c.hour%24, c.minute/60, c.minute%60
	if c.hour < 0 {
		x = 24 + x
	}
	if c.minute < 0 {
		if z != 0 {
			y = y - 1
			z = 60 + z
		}

	}

	c.hour = (x + y) % 24
	if c.hour < 0 {
		c.hour = 24 + c.hour
	}
	c.minute = z
	return c
}

//New ...
func New(hour, minute int) Clock {
	c := Clock{hour, minute}
	return c.clock()
}

//String ...
func (c Clock) String() string {
	hours := strconv.Itoa(c.hour)     //hour as a string
	minutes := strconv.Itoa(c.minute) //minute as a string
	if len(hours) != 2 {
		hours = "0" + hours
	}
	if len(minutes) != 2 {
		minutes = "0" + minutes

	}
	return hours + ":" + minutes
}

//Add ...
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.clock()
}

//Subtract ...
func (c Clock) Subtract(minutes int) Clock {
	c.minute -= minutes
	return c.clock()
}
