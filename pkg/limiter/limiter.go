package limiter

import (
	"time"
)

// Token represents a token of limiter.
type Token struct {
	createAt time.Time
}

// NewToken creates a token instance.
func NewToken() Token {
	return Token{
		createAt: time.Now().UTC(),
	}
}

// Limiter is an interface which limits the way to request token.
type Limiter interface {
	RequestToken() (token Token, err error)
	Cancel()
}
