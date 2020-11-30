package entity

import (
	"time"
)

// RequestToken represents a token includs the index between the limiter duration.
type RequestToken struct {
	index    uint32
	createAt time.Time
}

// NewRequestToken creates an instance of the request token.
func NewRequestToken(index uint32) RequestToken {
	return RequestToken{
		index:    index,
		createAt: time.Now().UTC(),
	}
}

// Index returns the index of the token.
func (t RequestToken) Index() uint32 {
	return t.index
}
