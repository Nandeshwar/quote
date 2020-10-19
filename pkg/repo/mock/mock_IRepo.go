// Code generated by MockGen. DO NOT EDIT.
// Source: loginrepo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRepo is a mock of IRepo interface
type MockIRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMockRecorder
}

// MockIRepoMockRecorder is the mock recorder for MockIRepo
type MockIRepoMockRecorder struct {
	mock *MockIRepo
}

// NewMockIRepo creates a new mock instance
func NewMockIRepo(ctrl *gomock.Controller) *MockIRepo {
	mock := &MockIRepo{ctrl: ctrl}
	mock.recorder = &MockIRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRepo) EXPECT() *MockIRepoMockRecorder {
	return m.recorder
}

// LoginInfo mocks base method
func (m *MockIRepo) LoginInfo(user, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginInfo", user, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginInfo indicates an expected call of LoginInfo
func (mr *MockIRepoMockRecorder) LoginInfo(user, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginInfo", reflect.TypeOf((*MockIRepo)(nil).LoginInfo), user, password)
}
