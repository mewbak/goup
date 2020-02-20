// Code generated by MockGen. DO NOT EDIT.
// Source: module.go

// Package mock_mod is a generated GoMock package.
package mock_mod

import (
	gomock "github.com/golang/mock/gomock"
	semver "github.com/rvflash/goup/internal/semver"
	reflect "reflect"
)

// MockModule is a mock of Module interface
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// Indirect mocks base method
func (m *MockModule) Indirect() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Indirect")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Indirect indicates an expected call of Indirect
func (mr *MockModuleMockRecorder) Indirect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indirect", reflect.TypeOf((*MockModule)(nil).Indirect))
}

// Path mocks base method
func (m *MockModule) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockModuleMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockModule)(nil).Path))
}

// Version mocks base method
func (m *MockModule) Version() semver.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(semver.Tag)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockModuleMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockModule)(nil).Version))
}