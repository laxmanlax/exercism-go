package clock

import "fmt"

//TestVersion is the unit tests this is guarenteed to pass.
const TestVersion = 2

//Clock keeps time, limited to a day
type Clock struct {
	minutes int
}

/*Time creates a Clock set to a given time.*/
func Time(hour, minute int) Clock {
	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock{time}
}

/*String returns a clock in digital form hh:mm.*/
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

/*Add move the time by a number of minutes.*/
func (c Clock) Add(minutes int) Clock {
	time := (c.minutes + minutes) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock{time}
}