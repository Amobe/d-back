package repository

import (
	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/util/limiter"
)

// IPLimiterRepository represents the interface of access the data sources of the ip limiter service.
type IPLimiterRepository interface {
	// Save saves the token limiter with an ip address as the index.
	Save(ipAddress entity.IPAddress, tokenLimiter limiter.Limiter) error
	// GetByIP selects the token limiter with an ip address as the index.
	GetByIP(ipAddress entity.IPAddress) (limiter.Limiter, error)
}
