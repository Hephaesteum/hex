package backoff

import (
	"testing"
	"time"
)

func TestRandomBackoff(t *testing.T) {
	backoff := &RandomBackoff{}

	for i := 0; i < 10; i++ {
		duration := backoff.Duration(i)
		if duration < time.Duration(500*i)*time.Millisecond || duration > time.Duration(1500*i)*time.Millisecond {
			t.Errorf("unexpected backoff duration: %v", duration)
		}
	}
}

func BenchmarkRandomBackoff_Duration(b *testing.B) {
	// Initialize backoff policy
	backoff := &RandomBackoff{}

	// Run benchmark for increasing values of n
	for n := 0; n < b.N; n++ {
		backoff.Duration(n)
	}
}
