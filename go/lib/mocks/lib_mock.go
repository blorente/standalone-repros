// Code generated by MockGen. DO NOT EDIT.
// Source: go/lib/lib.go

// Package mock_lib is a generated GoMock package.
package mock_lib

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMyInterface is a mock of MyInterface interface.
type MockMyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMyInterfaceMockRecorder
}

// MockMyInterfaceMockRecorder is the mock recorder for MockMyInterface.
type MockMyInterfaceMockRecorder struct {
	mock *MockMyInterface
}

// NewMockMyInterface creates a new mock instance.
func NewMockMyInterface(ctrl *gomock.Controller) *MockMyInterface {
	mock := &MockMyInterface{ctrl: ctrl}
	mock.recorder = &MockMyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyInterface) EXPECT() *MockMyInterfaceMockRecorder {
	return m.recorder
}

// World mocks base method.
func (m *MockMyInterface) World(hello string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "World", hello)
	ret0, _ := ret[0].(string)
	return ret0
}

// World indicates an expected call of World.
func (mr *MockMyInterfaceMockRecorder) World(hello interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "World", reflect.TypeOf((*MockMyInterface)(nil).World), hello)
}
