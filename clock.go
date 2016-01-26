package clock

import "time"

// Clock provides access to the current local time.
type Clock interface {
	Now() time.Time
}

var instance = &clock{}

func Default() Clock {
	return instance
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type clock struct {
}

func (c *clock) Now() time.Time {
	return time.Now()
}
