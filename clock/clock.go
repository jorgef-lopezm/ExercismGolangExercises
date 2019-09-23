package clock

import "fmt"

const minutesInADay = 1440

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	minutes := (h * 60 + m) % minutesInADay

	if minutes < 0 {
		minutes += minutesInADay
	}

	return Clock{minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes / 60, c.minutes % 60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}