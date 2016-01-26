package clock

import "time"

// In provides access to the current time from the given clock in the location
// provided.
func In(c Clock, loc *time.Location) Clock {
	return ClockFunc(func() time.Time {
		return c.Now().In(loc)
	})
}
