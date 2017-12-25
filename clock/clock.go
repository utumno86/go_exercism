package clock

import "fmt"

const testVersion = 4

// Clock object
type Clock int

// New Clock
func New(hour, minute int) Clock {
	totalMinutes := (hour*60 + minute)
	for totalMinutes >= 1440 {
		totalMinutes -= 1440
	}
	for totalMinutes < 0 {
		totalMinutes += 1440
	}
	return Clock(totalMinutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", int(c)/60, int(c)%60)
}

//Add Adds minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}
