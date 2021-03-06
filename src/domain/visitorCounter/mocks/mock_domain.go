// Code generated by MockGen. DO NOT EDIT.
// Source: domain.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// IncreaseVisits mocks base method.
func (m *MockDomain) IncreaseVisits(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseVisits", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseVisits indicates an expected call of IncreaseVisits.
func (mr *MockDomainMockRecorder) IncreaseVisits(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseVisits", reflect.TypeOf((*MockDomain)(nil).IncreaseVisits), ctx)
}
