package limiter

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/amobe/d-back/pkg/entity"
)

type AcceptanceRateLimiter struct {
	acceptanceCount uint32
	limitNumber     uint32
	inDuration      time.Duration
	clearSignal     chan struct{}
	cancelSignal    chan struct{}
	cancelOnce      sync.Once
}

var _ Limiter = &AcceptanceRateLimiter{}

func NewAcceptanceRateLimiter(limitNumber uint32, inDuration time.Duration) Limiter {
	l := &AcceptanceRateLimiter{
		acceptanceCount: 0,
		limitNumber:     limitNumber,
		inDuration:      inDuration,
		clearSignal:     make(chan struct{}, 1),
		cancelSignal:    make(chan struct{}),
	}
	return l
}

func (l *AcceptanceRateLimiter) Accept() (entity.RequestToken, error) {
	if l.isCancel() {
		return entity.RequestToken{}, fmt.Errorf("limiter is closed")
	}
	l.clearTrigger()
	newCount := atomic.AddUint32(&l.acceptanceCount, 1)
	if newCount > l.limitNumber {
		return entity.RequestToken{}, fmt.Errorf("accpectance denied")
	}
	return entity.NewRequestToken(newCount), nil
}

func (l *AcceptanceRateLimiter) Cancel() {
	l.cancelOnce.Do(func() {
		close(l.clearSignal)
		close(l.cancelSignal)
	})
}

func (l *AcceptanceRateLimiter) isCancel() bool {
	select {
	case <-l.cancelSignal:
		return true
	default:
		return false
	}
}

func (l *AcceptanceRateLimiter) clearTrigger() {
	select {
	case l.clearSignal <- struct{}{}:
		go l.clear()
	default:
	}
}

func (l *AcceptanceRateLimiter) clear() {
	ticker := time.NewTicker(l.inDuration)
	select {
	case <-ticker.C:
		atomic.StoreUint32(&l.acceptanceCount, 0)
	}
	<-l.clearSignal
}
