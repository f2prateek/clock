package clock

import "time"

// Clock provides access to the current local time.
type Clock interface {
	Now() time.Time
}

// Default returns a clock that provides the current local time.
func Default() Clock {
	return ClockFunc(time.Now)
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}
