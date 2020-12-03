package limiter_test

import (
	"testing"
	"time"

	"github.com/amobe/d-back/pkg/util/limiter"
)

// BenchmarkAcceptanceRateLimiterRequestToken-6   	12931077	        95.8 ns/op	       0 B/op	       0 allocs/op
func BenchmarkAcceptanceRateLimiterRequestToken(b *testing.B) {
	// given a acceptance rate limit allow to accept 1 billion requests in 10 second.
	const (
		limitNumber = 1000000000
		inDuration  = time.Second * 10
	)
	l := limiter.NewAcceptanceRateLimiter(limitNumber, inDuration)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Accept()
	}
}
