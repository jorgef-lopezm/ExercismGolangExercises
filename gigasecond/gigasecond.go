package gigasecond

import "time"

// AddGigasecond returns the time of t + 10^9
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
