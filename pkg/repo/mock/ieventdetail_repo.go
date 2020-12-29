// Code generated by MockGen. DO NOT EDIT.
// Source: event-detail-repo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	gomock "github.com/golang/mock/gomock"
	model "quote/pkg/model"
	reflect "reflect"
)

// MockIEventDetailRepo is a mock of IEventDetailRepo interface
type MockIEventDetailRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIEventDetailRepoMockRecorder
}

// MockIEventDetailRepoMockRecorder is the mock recorder for MockIEventDetailRepo
type MockIEventDetailRepoMockRecorder struct {
	mock *MockIEventDetailRepo
}

// NewMockIEventDetailRepo creates a new mock instance
func NewMockIEventDetailRepo(ctrl *gomock.Controller) *MockIEventDetailRepo {
	mock := &MockIEventDetailRepo{ctrl: ctrl}
	mock.recorder = &MockIEventDetailRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventDetailRepo) EXPECT() *MockIEventDetailRepoMockRecorder {
	return m.recorder
}

// CreateEventDetail mocks base method
func (m *MockIEventDetailRepo) CreateEventDetail(eventDetail model.EventDetail) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventDetail", eventDetail)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEventDetail indicates an expected call of CreateEventDetail
func (mr *MockIEventDetailRepoMockRecorder) CreateEventDetail(eventDetail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventDetail", reflect.TypeOf((*MockIEventDetailRepo)(nil).CreateEventDetail), eventDetail)
}

// GetEventDetailByTitleOrInfo mocks base method
func (m *MockIEventDetailRepo) GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByTitleOrInfo", searchTxt)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByTitleOrInfo indicates an expected call of GetEventDetailByTitleOrInfo
func (mr *MockIEventDetailRepoMockRecorder) GetEventDetailByTitleOrInfo(searchTxt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByTitleOrInfo", reflect.TypeOf((*MockIEventDetailRepo)(nil).GetEventDetailByTitleOrInfo), searchTxt)
}

// GetEventDetailByMonthDay mocks base method
func (m *MockIEventDetailRepo) GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByMonthDay", month, day)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByMonthDay indicates an expected call of GetEventDetailByMonthDay
func (mr *MockIEventDetailRepoMockRecorder) GetEventDetailByMonthDay(month, day interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByMonthDay", reflect.TypeOf((*MockIEventDetailRepo)(nil).GetEventDetailByMonthDay), month, day)
}

// GetEventDetailByMonth mocks base method
func (m *MockIEventDetailRepo) GetEventDetailByMonth(month int) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByMonth", month)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByMonth indicates an expected call of GetEventDetailByMonth
func (mr *MockIEventDetailRepoMockRecorder) GetEventDetailByMonth(month interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByMonth", reflect.TypeOf((*MockIEventDetailRepo)(nil).GetEventDetailByMonth), month)
}

// GetEventDetailByID mocks base method
func (m *MockIEventDetailRepo) GetEventDetailByID(ID int64) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByID", ID)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByID indicates an expected call of GetEventDetailByID
func (mr *MockIEventDetailRepoMockRecorder) GetEventDetailByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByID", reflect.TypeOf((*MockIEventDetailRepo)(nil).GetEventDetailByID), ID)
}

// UpdateEventDetailByID mocks base method
func (m *MockIEventDetailRepo) UpdateEventDetailByID(eventDetail model.EventDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventDetailByID", eventDetail)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventDetailByID indicates an expected call of UpdateEventDetailByID
func (mr *MockIEventDetailRepoMockRecorder) UpdateEventDetailByID(eventDetail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventDetailByID", reflect.TypeOf((*MockIEventDetailRepo)(nil).UpdateEventDetailByID), eventDetail)
}

// GetEventLinkIDs mocks base method
func (m *MockIEventDetailRepo) GetEventLinkIDs(links []string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventLinkIDs", links)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventLinkIDs indicates an expected call of GetEventLinkIDs
func (mr *MockIEventDetailRepoMockRecorder) GetEventLinkIDs(links interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventLinkIDs", reflect.TypeOf((*MockIEventDetailRepo)(nil).GetEventLinkIDs), links)
}
