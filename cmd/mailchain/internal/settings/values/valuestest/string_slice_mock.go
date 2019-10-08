// Code generated by MockGen. DO NOT EDIT.
// Source: string_slice.go

// Package valuestest is a generated GoMock package.
package valuestest

import (
	gomock "github.com/golang/mock/gomock"
	output "github.com/mailchain/mailchain/cmd/mailchain/internal/settings/output"
	reflect "reflect"
)

// MockStringSlice is a mock of StringSlice interface
type MockStringSlice struct {
	ctrl     *gomock.Controller
	recorder *MockStringSliceMockRecorder
}

// MockStringSliceMockRecorder is the mock recorder for MockStringSlice
type MockStringSliceMockRecorder struct {
	mock *MockStringSlice
}

// NewMockStringSlice creates a new mock instance
func NewMockStringSlice(ctrl *gomock.Controller) *MockStringSlice {
	mock := &MockStringSlice{ctrl: ctrl}
	mock.recorder = &MockStringSliceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStringSlice) EXPECT() *MockStringSliceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockStringSlice) Get() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockStringSliceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStringSlice)(nil).Get))
}

// Set mocks base method
func (m *MockStringSlice) Set(v []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", v)
}

// Set indicates an expected call of Set
func (mr *MockStringSliceMockRecorder) Set(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStringSlice)(nil).Set), v)
}

// Attribute mocks base method
func (m *MockStringSlice) Attribute() output.Attribute {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attribute")
	ret0, _ := ret[0].(output.Attribute)
	return ret0
}

// Attribute indicates an expected call of Attribute
func (mr *MockStringSliceMockRecorder) Attribute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attribute", reflect.TypeOf((*MockStringSlice)(nil).Attribute))
}
