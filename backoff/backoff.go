package backoff

import (
	"math/rand"
	"time"
)

func init() {
	// Seed random number generator
	rand.Seed(time.Now().UTC().UnixNano())
}

// Backoff implemented by Policies.
type Backoff interface {
	Duration(n int) time.Duration
}

type RandomBackoff struct{}

// Duration returns randomized delay between 0.5 * millis and 1.5 * millis.
func (b *RandomBackoff) Duration(n int) time.Duration {
	delayMillis := rand.Intn(1000 * (1 << n))
	jitterMillis := jitter(delayMillis)
	return time.Duration(delayMillis+jitterMillis) * time.Millisecond
}

// jitter returns random integer uniformly distributed in the range
// [0.5 * millis .. 1.5 * millis]
func jitter(millis int) int {
	if millis == 0 {
		return 0
	}

	return (millis / 2) + rand.Intn(millis/2+1)
}
