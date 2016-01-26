package clock

import "time"

// Clock provides access to the current time.
type Clock interface {
	Now() time.Time
}

// Default clock provides access to the current time in the local time zone.
func Default() Clock {
	return ClockFunc(time.Now)
}

// DefaultIn provides access to the current time from the default clock in the
// location provided.
func DefaultIn(loc *time.Location) Clock {
	return In(Default(), loc)
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}
