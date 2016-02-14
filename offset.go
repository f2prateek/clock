package clock

import "time"

// Offset returns a clock that always returns a time from the given clock with
// the specified duration added.
func Offset(c Clock, offset time.Duration) Clock {
	return ClockFunc(func() time.Time {
		return c.Now().Add(offset)
	})
}
