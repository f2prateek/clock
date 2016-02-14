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

// The ClockFunc type is an adapter to allow the use of ordinary functions as
// a Clock. If f is a function with the appropriate signature,
// ClockFunc(f) is a Clock object that calls f.
type ClockFunc func() time.Time

// Now calls f().
func (f ClockFunc) Now() time.Time {
	return f()
}
