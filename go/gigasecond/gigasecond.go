package gigasecond

import "time"

// Adds 1_000_000_000 (10 ^ 9) seconds to the time input
func AddGigasecond(t time.Time) time.Time {
	const gs = 1_000_000_000
	return t.Add(gs * time.Second)
}
