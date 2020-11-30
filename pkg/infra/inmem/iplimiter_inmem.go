package inmem

import (
	"fmt"
	"sync"

	"github.com/amobe/d-back/pkg/exception"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/repository"
	"github.com/amobe/d-back/pkg/util/limiter"
)

type ipLimiterInMem struct {
	storage sync.Map
}

var _ repository.IPLimiterRepository = &ipLimiterInMem{}

// NewIPLimiterRepository creates a ip limiter repository which implements with local memory.
func NewIPLimiterRepository() repository.IPLimiterRepository {
	return &ipLimiterInMem{}
}

// Save saves a limiter with ip address.
func (r *ipLimiterInMem) Save(ipAddress entity.IPAddress, tokenLimiter limiter.Limiter) error {
	r.storage.Store(ipAddress.Host(), tokenLimiter)
	return nil
}

// GetByIP returns a limiter by ip address.
func (r *ipLimiterInMem) GetByIP(ipAddress entity.IPAddress) (limiter.Limiter, error) {
	data, ok := r.storage.Load(ipAddress.Host())
	if !ok {
		return nil, exception.ErrRepositoryNotFound
	}
	tokenLimiter, ok := data.(limiter.Limiter)
	if !ok {
		return nil, fmt.Errorf("invalid data type")
	}
	return tokenLimiter, nil
}
