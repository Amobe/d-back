package limiter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/amobe/d-back/pkg/limiter"
	"github.com/stretchr/testify/assert"
)

func TestThrottleLimited(t *testing.T) {
	// given a throttle rate limit allow to accept 2 request in 1 second.
	cfg := limiter.ThrottleConfig{
		LimitDuration: time.Second,
		LimitTimes:    2,
	}
	l := limiter.NewThrottleRateLimiter(cfg)

	// when request 3 token in 1 second
	const requestTimes = 3
	errMap := make(map[int]error, requestTimes)
	for i := 0; i < requestTimes; i++ {
		_, err := l.RequestToken()
		errMap[i] = err
	}

	// then the first and second request should be accpected
	assert.NoError(t, errMap[0])
	assert.NoError(t, errMap[1])

	// and the third request should be denied
	assert.Error(t, errMap[2])

	l.Cancel()
}

func TestThrottleRecovered(t *testing.T) {
	requestDuration := time.Millisecond
	extraDuration := time.Millisecond

	// given a throttle rate limit allow to accept 2 request in 1 second.
	cfg := limiter.ThrottleConfig{
		LimitDuration: requestDuration,
		LimitTimes:    2,
	}
	l := limiter.NewThrottleRateLimiter(cfg)

	// when request 1 token
	// and wait 1 second
	// and request 2 tokens
	errMap := make(map[int]error, 3)
	_, err := l.RequestToken()
	errMap[0] = err
	// FIXME: between request and reload need a little bit extraDuration
	time.Sleep(requestDuration + extraDuration)
	_, err = l.RequestToken()
	errMap[1] = err
	_, err = l.RequestToken()
	errMap[2] = err

	// then all requests should be accepted
	assert.NoError(t, errMap[0])
	assert.NoError(t, errMap[1])
	assert.NoError(t, errMap[2])

	l.Cancel()
}

func TestThrottleConcurrency(t *testing.T) {
	const requestTimes = 100

	// given a throttle rate limit allow to accept 100 request in 1 second.
	cfg := limiter.ThrottleConfig{
		LimitDuration: time.Second,
		LimitTimes:    requestTimes,
	}
	l := limiter.NewThrottleRateLimiter(cfg)

	// when request 100 tokens concurrency
	var wg sync.WaitGroup
	wg.Add(requestTimes)
	errCh := make(chan error, requestTimes)
	for i := 0; i < requestTimes; i++ {
		go func() {
			_, err := l.RequestToken()
			errCh <- err
			wg.Done()
		}()
	}
	wg.Wait()
	close(errCh)

	// then the all request should be accpected
	for err := range errCh {
		assert.NoError(t, err)
	}

	l.Cancel()
}
