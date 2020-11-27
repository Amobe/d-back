package limiter

import (
	"fmt"
	"sync"
	"time"
)

// ThrottleConfig represnts a throttle config include limited information.
type ThrottleConfig struct {
	LimitDuration time.Duration
	LimitTimes    int
}

// ThrottleRateLimiter limits the requests in duration.
type ThrottleRateLimiter struct {
	cfg          ThrottleConfig
	requestChan  chan struct{}
	tokenChan    chan Token
	cancelSignal chan struct{}
	cancelOnce   sync.Once
}

// ThrottleRateLimiter implements the RateLimiter interface.
var _ Limiter = &ThrottleRateLimiter{}

// NewThrottleRateLimiter creates the ThrottleRateLimiter instance.
func NewThrottleRateLimiter(cfg ThrottleConfig) Limiter {
	l := &ThrottleRateLimiter{
		cfg:          cfg,
		requestChan:  make(chan struct{}, 1),
		tokenChan:    make(chan Token, cfg.LimitTimes),
		cancelSignal: make(chan struct{}),
	}
	l.reloadToken()
	go l.await()
	return l
}

// RequestToken returns the token when request is accepted.
// Retrun err when requests reach the limitation.
func (t *ThrottleRateLimiter) RequestToken() (token Token, err error) {
	select {
	case t.requestChan <- struct{}{}:
	default:
	}
	select {
	case token := <-t.tokenChan:
		return token, nil
	default:
		return Token{}, fmt.Errorf("request denied")
	}
}

// Cancel release the resource use by the limiter
func (t *ThrottleRateLimiter) Cancel() {
	t.cancelOnce.Do(func() {
		close(t.requestChan)
		close(t.tokenChan)
		close(t.cancelSignal)
	})
}

func (t *ThrottleRateLimiter) reloadToken() {
	for i := 0; i < t.cfg.LimitTimes; i++ {
		select {
		case t.tokenChan <- NewToken():
		default:
			// exit when channel is full
			return
		}
	}
}

func (t *ThrottleRateLimiter) await() {
	timer := time.NewTimer(t.cfg.LimitDuration)
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}
	for range t.requestChan {
		timer.Reset(t.cfg.LimitDuration)
		select {
		case <-timer.C:
			t.reloadToken()
		case <-t.cancelSignal:
			timer.Stop()
			return
		}
	}
}
