// Code generated by MockGen. DO NOT EDIT.
// Source: inforepo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "quote/pkg/model"
	reflect "reflect"
)

// MockIInfoRepo is a mock of IInfoRepo interface
type MockIInfoRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIInfoRepoMockRecorder
}

// MockIInfoRepoMockRecorder is the mock recorder for MockIInfoRepo
type MockIInfoRepoMockRecorder struct {
	mock *MockIInfoRepo
}

// NewMockIInfoRepo creates a new mock instance
func NewMockIInfoRepo(ctrl *gomock.Controller) *MockIInfoRepo {
	mock := &MockIInfoRepo{ctrl: ctrl}
	mock.recorder = &MockIInfoRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIInfoRepo) EXPECT() *MockIInfoRepoMockRecorder {
	return m.recorder
}

// CreateInfo mocks base method
func (m *MockIInfoRepo) CreateInfo(ctx context.Context, info model.Info) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInfo", ctx, info)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInfo indicates an expected call of CreateInfo
func (mr *MockIInfoRepoMockRecorder) CreateInfo(ctx, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInfo", reflect.TypeOf((*MockIInfoRepo)(nil).CreateInfo), ctx, info)
}

// GetInfoByTitleOrInfo mocks base method
func (m *MockIInfoRepo) GetInfoByTitleOrInfo(searchTxt string) ([]model.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoByTitleOrInfo", searchTxt)
	ret0, _ := ret[0].([]model.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoByTitleOrInfo indicates an expected call of GetInfoByTitleOrInfo
func (mr *MockIInfoRepoMockRecorder) GetInfoByTitleOrInfo(searchTxt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoByTitleOrInfo", reflect.TypeOf((*MockIInfoRepo)(nil).GetInfoByTitleOrInfo), searchTxt)
}

// UpdateInfoByID mocks base method
func (m *MockIInfoRepo) UpdateInfoByID(info model.Info) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfoByID", info)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfoByID indicates an expected call of UpdateInfoByID
func (mr *MockIInfoRepoMockRecorder) UpdateInfoByID(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfoByID", reflect.TypeOf((*MockIInfoRepo)(nil).UpdateInfoByID), info)
}

// GetInfoByID mocks base method
func (m *MockIInfoRepo) GetInfoByID(ctx context.Context, ID int64) ([]model.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoByID", ctx, ID)
	ret0, _ := ret[0].([]model.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoByID indicates an expected call of GetInfoByID
func (mr *MockIInfoRepoMockRecorder) GetInfoByID(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoByID", reflect.TypeOf((*MockIInfoRepo)(nil).GetInfoByID), ctx, ID)
}

// GetInfoLinkIDs mocks base method
func (m *MockIInfoRepo) GetInfoLinkIDs(links []string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfoLinkIDs", links)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfoLinkIDs indicates an expected call of GetInfoLinkIDs
func (mr *MockIInfoRepoMockRecorder) GetInfoLinkIDs(links interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfoLinkIDs", reflect.TypeOf((*MockIInfoRepo)(nil).GetInfoLinkIDs), links)
}
