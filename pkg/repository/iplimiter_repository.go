package repository

import (
	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/limiter"
)

type IPLimiterRepository interface {
	Save(ipAddress entity.IPAddress, tokenLimiter limiter.Limiter) error
	GetByIP(ipAddress entity.IPAddress) (limiter.Limiter, error)
}
