package limiter

import (
	"github.com/amobe/d-back/pkg/entity"
)

// Limiter is an interface which limits the way to request token.
type Limiter interface {
	// Accept returns the acceptance token if the request is not denied.
	Accept() (token entity.RequestToken, err error)
	// Cancel cancels the limiter to release the resource.
	Cancel()
}
