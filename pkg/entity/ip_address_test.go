package entity_test

import (
	"testing"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewIPAddress(t *testing.T) {
	const address = "127.0.0.1:65535"

	ipAddress, err := entity.NewIPAddress(address)
	assert.NoError(t, err)
	assert.Equal(t, "127.0.0.1", ipAddress.Host())
	assert.Equal(t, "65535", ipAddress.Port())
}

func TestNewIPAddressInvalidHost(t *testing.T) {
	const address = "....1:5678"

	_, err := entity.NewIPAddress(address)
	assert.Error(t, err)
}

func TestNewIPAddressInvalidPort(t *testing.T) {
	const address = "127.0.0.1:a"

	_, err := entity.NewIPAddress(address)
	assert.Error(t, err)
}
