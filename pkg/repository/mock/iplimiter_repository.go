// Code generated by MockGen. DO NOT EDIT.
// Source: iplimiter_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "github.com/amobe/d-back/pkg/entity"
	limiter "github.com/amobe/d-back/pkg/util/limiter"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPLimiterRepository is a mock of IPLimiterRepository interface
type MockIPLimiterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPLimiterRepositoryMockRecorder
}

// MockIPLimiterRepositoryMockRecorder is the mock recorder for MockIPLimiterRepository
type MockIPLimiterRepositoryMockRecorder struct {
	mock *MockIPLimiterRepository
}

// NewMockIPLimiterRepository creates a new mock instance
func NewMockIPLimiterRepository(ctrl *gomock.Controller) *MockIPLimiterRepository {
	mock := &MockIPLimiterRepository{ctrl: ctrl}
	mock.recorder = &MockIPLimiterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPLimiterRepository) EXPECT() *MockIPLimiterRepositoryMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockIPLimiterRepository) Save(ipAddress entity.IPAddress, tokenLimiter limiter.Limiter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ipAddress, tokenLimiter)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockIPLimiterRepositoryMockRecorder) Save(ipAddress, tokenLimiter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIPLimiterRepository)(nil).Save), ipAddress, tokenLimiter)
}

// GetByIP mocks base method
func (m *MockIPLimiterRepository) GetByIP(ipAddress entity.IPAddress) (limiter.Limiter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIP", ipAddress)
	ret0, _ := ret[0].(limiter.Limiter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIP indicates an expected call of GetByIP
func (mr *MockIPLimiterRepositoryMockRecorder) GetByIP(ipAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIP", reflect.TypeOf((*MockIPLimiterRepository)(nil).GetByIP), ipAddress)
}
