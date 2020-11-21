// Code generated by MockGen. DO NOT EDIT.
// Source: event-detail-service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	model "quote/pkg/model"
	reflect "reflect"
	time "time"
)

// MockIEventDetail is a mock of IEventDetail interface
type MockIEventDetail struct {
	ctrl     *gomock.Controller
	recorder *MockIEventDetailMockRecorder
}

// MockIEventDetailMockRecorder is the mock recorder for MockIEventDetail
type MockIEventDetailMockRecorder struct {
	mock *MockIEventDetail
}

// NewMockIEventDetail creates a new mock instance
func NewMockIEventDetail(ctrl *gomock.Controller) *MockIEventDetail {
	mock := &MockIEventDetail{ctrl: ctrl}
	mock.recorder = &MockIEventDetailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventDetail) EXPECT() *MockIEventDetailMockRecorder {
	return m.recorder
}

// ValidateFormEvent mocks base method
func (m *MockIEventDetail) ValidateFormEvent(form model.EventDetailForm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateFormEvent", form)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateFormEvent indicates an expected call of ValidateFormEvent
func (mr *MockIEventDetailMockRecorder) ValidateFormEvent(form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateFormEvent", reflect.TypeOf((*MockIEventDetail)(nil).ValidateFormEvent), form)
}

// CreateNewEventDetail mocks base method
func (m *MockIEventDetail) CreateNewEventDetail(form model.EventDetailForm) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewEventDetail", form)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewEventDetail indicates an expected call of CreateNewEventDetail
func (mr *MockIEventDetailMockRecorder) CreateNewEventDetail(form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewEventDetail", reflect.TypeOf((*MockIEventDetail)(nil).CreateNewEventDetail), form)
}

// GetEventDetailByTitleOrInfo mocks base method
func (m *MockIEventDetail) GetEventDetailByTitleOrInfo(searchTxt string) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByTitleOrInfo", searchTxt)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByTitleOrInfo indicates an expected call of GetEventDetailByTitleOrInfo
func (mr *MockIEventDetailMockRecorder) GetEventDetailByTitleOrInfo(searchTxt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByTitleOrInfo", reflect.TypeOf((*MockIEventDetail)(nil).GetEventDetailByTitleOrInfo), searchTxt)
}

// GetEventDetailByMonthDay mocks base method
func (m *MockIEventDetail) GetEventDetailByMonthDay(month, day int) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByMonthDay", month, day)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByMonthDay indicates an expected call of GetEventDetailByMonthDay
func (mr *MockIEventDetailMockRecorder) GetEventDetailByMonthDay(month, day interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByMonthDay", reflect.TypeOf((*MockIEventDetail)(nil).GetEventDetailByMonthDay), month, day)
}

// GetEventDetailByID mocks base method
func (m *MockIEventDetail) GetEventDetailByID(ID int64) (model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailByID", ID)
	ret0, _ := ret[0].(model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailByID indicates an expected call of GetEventDetailByID
func (mr *MockIEventDetailMockRecorder) GetEventDetailByID(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailByID", reflect.TypeOf((*MockIEventDetail)(nil).GetEventDetailByID), ID)
}

// UpdateEventDetailByID mocks base method
func (m *MockIEventDetail) UpdateEventDetailByID(eventDetail model.EventDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEventDetailByID", eventDetail)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEventDetailByID indicates an expected call of UpdateEventDetailByID
func (mr *MockIEventDetailMockRecorder) UpdateEventDetailByID(eventDetail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEventDetailByID", reflect.TypeOf((*MockIEventDetail)(nil).UpdateEventDetailByID), eventDetail)
}

// GetEventDetailLinkIDs mocks base method
func (m *MockIEventDetail) GetEventDetailLinkIDs(link string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventDetailLinkIDs", link)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventDetailLinkIDs indicates an expected call of GetEventDetailLinkIDs
func (mr *MockIEventDetailMockRecorder) GetEventDetailLinkIDs(link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventDetailLinkIDs", reflect.TypeOf((*MockIEventDetail)(nil).GetEventDetailLinkIDs), link)
}

// EventsInFuture mocks base method
func (m *MockIEventDetail) EventsInFuture(t time.Time) ([]model.EventDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsInFuture", t)
	ret0, _ := ret[0].([]model.EventDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsInFuture indicates an expected call of EventsInFuture
func (mr *MockIEventDetailMockRecorder) EventsInFuture(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsInFuture", reflect.TypeOf((*MockIEventDetail)(nil).EventsInFuture), t)
}
