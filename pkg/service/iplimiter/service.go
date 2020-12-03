package iplimiter

import (
	"errors"
	"fmt"
	"time"

	"github.com/amobe/d-back/pkg/exception"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/repository"
	"github.com/amobe/d-back/pkg/util/limiter"
)

// Service represents the ip limiter service.
type Service interface {
	AcceptRequest(ipAddress entity.IPAddress) (entity.RequestToken, error)
}

type service struct {
	ipLimiterRepository repository.IPLimiterRepository

	ipRequestLimitNumber   uint32
	ipRequestLimitDuration time.Duration
}

var _ Service = &service{}

// NewIPLimiterService creates the instance of ip limiter service.
func NewIPLimiterService(ipLimiterRepository repository.IPLimiterRepository, limitNumber uint32, limitDuration time.Duration) Service {
	return &service{
		ipLimiterRepository:    ipLimiterRepository,
		ipRequestLimitNumber:   limitNumber,
		ipRequestLimitDuration: limitDuration,
	}
}

// AcceptRequest accepts the request and returns token for specific ip address.
func (s *service) AcceptRequest(ipAddress entity.IPAddress) (entity.RequestToken, error) {
	tokenLimiter, err := s.getOrCreateTokenLimiter(ipAddress)
	if err != nil {
		return entity.RequestToken{}, fmt.Errorf("get or create token limiter: %w", err)
	}
	token, err := tokenLimiter.Accept()
	if err != nil {
		return entity.RequestToken{}, fmt.Errorf("request token: %w", err)
	}
	return token, nil
}

func (s *service) getOrCreateTokenLimiter(ipAddress entity.IPAddress) (limiter.Limiter, error) {
	tokenLimiter, err := s.ipLimiterRepository.GetByIP(ipAddress)
	if err == nil {
		return tokenLimiter, nil
	}
	if !errors.Is(err, exception.ErrRepositoryNotFound) {
		return nil, fmt.Errorf("get token limiter: %w", err)
	}

	tokenLimiter = limiter.NewAcceptanceRateLimiter(s.ipRequestLimitNumber, s.ipRequestLimitDuration)
	if err := s.ipLimiterRepository.Save(ipAddress, tokenLimiter); err != nil {
		return nil, fmt.Errorf("save token limiter: %w", err)
	}
	return tokenLimiter, nil
}
