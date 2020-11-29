package iplimiter

import (
	"errors"
	"fmt"
	"time"

	"github.com/amobe/d-back/pkg/exception"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/limiter"
	"github.com/amobe/d-back/pkg/repository"
)

// Service represents the ip limiter service.
type Service interface {
	AcceptRequest(ipAddress entity.IPAddress) (limiter.Token, error)
}

type service struct {
	ipLimiterRepository repository.IPLimiterRepository
}

var _ Service = &service{}

// NewIPLimiterService creates the instance of ip limiter service.
func NewIPLimiterService(ipLimiterRepository repository.IPLimiterRepository) Service {
	return &service{
		ipLimiterRepository: ipLimiterRepository,
	}
}

// AcceptRequest accepts the request and returns token for specific ip address.
func (s *service) AcceptRequest(ipAddress entity.IPAddress) (limiter.Token, error) {
	throttleLimiter, err := s.getOrCreateThrottleLimiter(ipAddress)
	if err != nil {
		return limiter.Token{}, fmt.Errorf("get or create throttle limiter: %w", err)
	}
	token, err := throttleLimiter.RequestToken()
	if err != nil {
		return limiter.Token{}, fmt.Errorf("request token: %w", err)
	}
	return token, nil
}

func (s *service) getOrCreateThrottleLimiter(ipAddress entity.IPAddress) (limiter.Limiter, error) {
	throttleLimiter, err := s.ipLimiterRepository.GetByIP(ipAddress)
	if err == nil {
		return throttleLimiter, nil
	}
	if !errors.Is(err, exception.ErrRepositoryNotFound) {
		return nil, fmt.Errorf("get throttle: %w", err)
	}

	throttleLimiter = limiter.NewThrottleRateLimiter(limiter.ThrottleConfig{
		LimitDuration: time.Minute,
		LimitTimes:    60,
	})
	if err := s.ipLimiterRepository.Save(ipAddress, throttleLimiter); err != nil {
		return nil, fmt.Errorf("save throttle limiter: %w", err)
	}
	return throttleLimiter, nil
}
