package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

const testVersion = 4

// AddGigasecond Adds a gigasecond to a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(int(math.Pow(10, 9))))
}
