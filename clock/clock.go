package clock

import (
	"fmt"
)

const (
	minInHour = 60
	minInDay  = 1440
)

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	totalMinutes := ((h * minInHour) + m) % minInDay

	if totalMinutes < 0 {
		totalMinutes = minInDay + totalMinutes
	}

	hour := totalMinutes / minInHour
	minute := totalMinutes % minInHour

	return Clock{
		hour,
		minute,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
