//Package clock ..
package clock

import "strconv"

//Clock ...
type Clock struct {
    hour, minute int
}

//New ...
func New(hour, minute int) Clock  {
    var c Clock
    // initialise and Add
    hour, minute = hour%24, minute%24
    if hour < 0 {
        hour = 24 - hour
        if minute < 0 {
            minute = 24 - minute
        }
    }
    c.hour = hour
    c.minute = minute
    return c
    }

//String ...
func (c Clock) String() string  {
    return strconv.Itoa(c.hour/10) + strconv.Itoa(c.hour%10) + // hours
":" +                                                         // :
    strconv.Itoa(c.minute/10) + strconv.Itoa(c.minute%10)      // minutes
}

//Add ...
func (c Clock) Add(minutes int) Clock {
    minutes = (c.minute + minutes) % 60
    hour := minutes / 60
    if minutes < 0 {
        minutes = 60 - minutes
        if hour < 0 {
            hour*= -1
        }
    }
    c.hour = c.hour + hour
    c.minute = minutes
    return c
}

//Subtract ...
func (c Clock) Subtract(minutes int) Clock {
    minutes*= -1
    return c.Add(minutes)
}
