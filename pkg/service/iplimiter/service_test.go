package iplimiter_test

import (
	"testing"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/exception"
	"github.com/amobe/d-back/pkg/limiter"
	mock_limiter "github.com/amobe/d-back/pkg/limiter/mock"
	mock_repository "github.com/amobe/d-back/pkg/repository/mock"
	"github.com/amobe/d-back/pkg/service/iplimiter"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceAcceptRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given an ip address has requested in the past
	// and the matched limiter already stored in repository
	ipAddress, err := entity.NewIPAddress("127.0.0.1:5678")
	assert.NoError(t, err)

	mockLimiter := mock_limiter.NewMockLimiter(ctrl)
	mockLimiter.EXPECT().RequestToken().Return(limiter.NewToken(), nil)

	ipLimiterRepository := mock_repository.NewMockIPLimiterRepository(ctrl)
	ipLimiterRepository.EXPECT().GetByIP(ipAddress).Return(mockLimiter, nil)

	ipLimiterService := iplimiter.NewIPLimiterService(ipLimiterRepository)

	// when the ip address requests again
	_, err = ipLimiterService.AcceptRequest(ipAddress)

	// then the request should be accept
	assert.NoError(t, err)
}

func TestServiceAcceptRequestFirstTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given an ip address has never requested in the past
	// and the matched limiter is not stored in repository
	ipAddress, err := entity.NewIPAddress("127.0.0.1:5678")
	assert.NoError(t, err)

	ipLimiterRepository := mock_repository.NewMockIPLimiterRepository(ctrl)
	ipLimiterRepository.EXPECT().GetByIP(ipAddress).Return(nil, exception.ErrRepositoryNotFound)
	ipLimiterRepository.EXPECT().Save(ipAddress, gomock.Not(gomock.Nil())).Return(nil)

	ipLimiterService := iplimiter.NewIPLimiterService(ipLimiterRepository)

	// when the ip address requests
	_, err = ipLimiterService.AcceptRequest(ipAddress)

	// then the a limiter should be created
	// and the limiter should be save to repository
	// and the requset should be accept
	assert.NoError(t, err)
}
