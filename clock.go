package clock

import "time"

// Clock provides access to the current time.
type Clock interface {
	Now() time.Time
}

// Clock provides access to the current time in the local time zone.
func Default() Clock {
	return ClockFunc(time.Now)
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}
