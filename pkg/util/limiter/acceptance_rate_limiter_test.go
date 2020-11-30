package limiter_test

import (
	"sync"
	"testing"
	"time"

	"github.com/amobe/d-back/pkg/entity"

	"github.com/amobe/d-back/pkg/limiter"
	"github.com/stretchr/testify/assert"
)

func TestAcceptanceRateLimiterLimited(t *testing.T) {
	// given a acceptance rate limit allow to accept 2 request in 1 second.
	const (
		limitNumber = 2
		inDuration  = time.Second
	)
	l := limiter.NewAcceptanceRateLimiter(limitNumber, inDuration)

	// when request 3 token in 1 second
	const requestTimes = 3
	errMap := make(map[int]error, requestTimes)
	for i := 0; i < requestTimes; i++ {
		_, err := l.Accept()
		errMap[i] = err
	}

	// then the first and second request should be accpected
	assert.NoError(t, errMap[0])
	assert.NoError(t, errMap[1])

	// and the third request should be denied
	assert.Error(t, errMap[2])

	l.Cancel()
}

func TestAcceptanceRateLimiterRecovered(t *testing.T) {
	waitRecovered := time.Millisecond * 2

	// given a acceptance rate limit allow to accept 2 request in 1 second.
	const (
		limitNumber = 2
		inDuration  = time.Millisecond
	)
	l := limiter.NewAcceptanceRateLimiter(limitNumber, inDuration)

	// when request 1 token
	// and wait 1 second
	// and request 2 tokens
	tokenMap := make(map[int]entity.RequestToken, 3)
	errMap := make(map[int]error, 3)
	token, err := l.Accept()
	tokenMap[0] = token
	errMap[0] = err
	time.Sleep(waitRecovered)
	token, err = l.Accept()
	tokenMap[1] = token
	errMap[1] = err
	token, err = l.Accept()
	tokenMap[2] = token
	errMap[2] = err

	// then all requests should be accepted
	assert.NoError(t, errMap[0])
	assert.Equal(t, tokenMap[0].Index(), uint32(1))
	assert.NoError(t, errMap[1])
	assert.Equal(t, tokenMap[1].Index(), uint32(1))
	assert.NoError(t, errMap[2])
	assert.Equal(t, tokenMap[2].Index(), uint32(2))

	l.Cancel()
}

func TestAcceptanceRateLimiterConcurrency(t *testing.T) {
	const requestTimes = 100

	// given a acceptance rate limit allow to accept 100 request in 1 second.
	const (
		limitNumber = requestTimes
		inDuration  = time.Second
	)
	l := limiter.NewAcceptanceRateLimiter(limitNumber, inDuration)

	// when request 100 tokens concurrency
	var wg sync.WaitGroup
	wg.Add(requestTimes)
	errCh := make(chan error, requestTimes)
	for i := 0; i < requestTimes; i++ {
		go func() {
			_, err := l.Accept()
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

func TestAcceptanceRateLimiterAfterCancel(t *testing.T) {
	const requestTimes = 10

	// given a acceptance rate limit is cancelled.
	const (
		limitNumber = 1
		inDuration  = time.Second
	)
	l := limiter.NewAcceptanceRateLimiter(limitNumber, inDuration)
	l.Cancel()

	// when request multiple times
	errCh := make(chan error, requestTimes)
	for i := 0; i < requestTimes; i++ {
		_, err := l.Accept()
		errCh <- err
	}
	close(errCh)

	// then the all request should return error
	for err := range errCh {
		assert.Error(t, err)
	}
}
