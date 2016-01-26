package clock

import "time"

// Fixed returns a clock that always returns the given "fixed" time.
func Fixed(t time.Time) Clock {
	return ClockFunc(func() time.Time {
		return t
	})
}
