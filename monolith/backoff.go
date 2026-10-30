package monolith

import "time"

// Backoff returns the delay before retry attempt n: 100ms doubled per
// attempt, capped at 5s. A negative attempt counts as attempt 0.
func Backoff(attempt int) time.Duration {
	if attempt >= 6 {
		return 5 * time.Second
	}
	if attempt < 0 {
		attempt = 0
	}
	return 100 * time.Millisecond << attempt
}
