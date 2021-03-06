// Code generated by MockGen. DO NOT EDIT.
// Source: src/config/config.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// GetMaxPrintCount mocks base method.
func (m *MockConfig) GetMaxPrintCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxPrintCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxPrintCount indicates an expected call of GetMaxPrintCount.
func (mr *MockConfigMockRecorder) GetMaxPrintCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxPrintCount", reflect.TypeOf((*MockConfig)(nil).GetMaxPrintCount))
}
