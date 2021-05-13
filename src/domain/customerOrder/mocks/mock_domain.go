// Code generated by MockGen. DO NOT EDIT.
// Source: domain.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	coModels "gitlab.com/altiano/goreen-tea/src/domain/customerOrder/models"
)

// MockDomain is a mock of Domain interface.
type MockDomain struct {
	ctrl     *gomock.Controller
	recorder *MockDomainMockRecorder
}

// MockDomainMockRecorder is the mock recorder for MockDomain.
type MockDomainMockRecorder struct {
	mock *MockDomain
}

// NewMockDomain creates a new mock instance.
func NewMockDomain(ctrl *gomock.Controller) *MockDomain {
	mock := &MockDomain{ctrl: ctrl}
	mock.recorder = &MockDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDomain) EXPECT() *MockDomainMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDomain) Create(ctx context.Context, req coModels.CreateReq) (coModels.CustomerOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, req)
	ret0, _ := ret[0].(coModels.CustomerOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDomainMockRecorder) Create(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDomain)(nil).Create), ctx, req)
}

// NotifyRedeemedCode mocks base method.
func (m *MockDomain) NotifyRedeemedCode(ctx context.Context, co coModels.CustomerOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyRedeemedCode", ctx, co)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyRedeemedCode indicates an expected call of NotifyRedeemedCode.
func (mr *MockDomainMockRecorder) NotifyRedeemedCode(ctx, co interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyRedeemedCode", reflect.TypeOf((*MockDomain)(nil).NotifyRedeemedCode), ctx, co)
}

// NotifyWaiterName mocks base method.
func (m *MockDomain) NotifyWaiterName(ctx context.Context, co coModels.CustomerOrder, waiterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyWaiterName", ctx, co, waiterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyWaiterName indicates an expected call of NotifyWaiterName.
func (mr *MockDomainMockRecorder) NotifyWaiterName(ctx, co, waiterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyWaiterName", reflect.TypeOf((*MockDomain)(nil).NotifyWaiterName), ctx, co, waiterName)
}
