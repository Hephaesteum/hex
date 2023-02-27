package backoff

import (
	"math/rand"
	"time"
)

// Policy is a backoff policy that implements randomizing delays and saturating
// the final value in milliseconds.
type Policy struct {
	Millis []int
}

// Default is the default backoff policy ranging up to 5 seconds.
var Default = Policy{
	Millis: []int{0, 10, 10, 100, 100, 500, 500, 3000, 3000, 5000},
}

// Duration returns the duration of the n'th wait cycle in a backoff policy.
// This is p.Millis[n], randomized to avoid thundering herds.
func (p *Policy) Duration(n int) time.Duration {
	// If  wait cycle exceeds length of backoff policy set it to the final value
	// in the policy to saturate the backoff time.
	if n >= len(p.Millis) {
		n = len(p.Millis) - 1
	}

	// Jitter backoff time by adding random duration between 0 and half value of
	// backoff time to avoid thundering herds.
	backoffTime := time.Duration(p.Millis[n]+rand.Intn(p.Millis[n]/2)) * time.Millisecond

	return backoffTime
}
