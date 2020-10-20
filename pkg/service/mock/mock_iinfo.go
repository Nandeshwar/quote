// Code generated by MockGen. DO NOT EDIT.
// Source: InfoService.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	model "quote/pkg/model"
	reflect "reflect"
)

// MockIInfo is a mock of IInfo interface
type MockIInfo struct {
	ctrl     *gomock.Controller
	recorder *MockIInfoMockRecorder
}

// MockIInfoMockRecorder is the mock recorder for MockIInfo
type MockIInfoMockRecorder struct {
	mock *MockIInfo
}

// NewMockIInfo creates a new mock instance
func NewMockIInfo(ctrl *gomock.Controller) *MockIInfo {
	mock := &MockIInfo{ctrl: ctrl}
	mock.recorder = &MockIInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIInfo) EXPECT() *MockIInfoMockRecorder {
	return m.recorder
}

// ValidateForm mocks base method
func (m *MockIInfo) ValidateForm(form model.InfoForm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateForm", form)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateForm indicates an expected call of ValidateForm
func (mr *MockIInfoMockRecorder) ValidateForm(form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateForm", reflect.TypeOf((*MockIInfo)(nil).ValidateForm), form)
}

// CreateNewInfo mocks base method
func (m *MockIInfo) CreateNewInfo(form model.InfoForm) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewInfo", form)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewInfo indicates an expected call of CreateNewInfo
func (mr *MockIInfoMockRecorder) CreateNewInfo(form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewInfo", reflect.TypeOf((*MockIInfo)(nil).CreateNewInfo), form)
}

// GetInfoByTitleOrInfo mocks base method
func (m *MockIInfo) GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoByTitleOrInfo", searchTxt)
	ret0, _ := ret[0].([]model.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoByTitleOrInfo indicates an expected call of GetInfoByTitleOrInfo
func (mr *MockIInfoMockRecorder) GetInfoByTitleOrInfo(searchTxt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoByTitleOrInfo", reflect.TypeOf((*MockIInfo)(nil).GetInfoByTitleOrInfo), searchTxt)
}