// Code generated by MockGen. DO NOT EDIT.
// Source: factory.go

// Package gcalauth is a generated GoMock package.
package gcalauth

import (
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
	reflect "reflect"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// FileStore mocks base method
func (m *MockFactory) FileStore() (Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileStore")
	ret0, _ := ret[0].(Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileStore indicates an expected call of FileStore
func (mr *MockFactoryMockRecorder) FileStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileStore", reflect.TypeOf((*MockFactory)(nil).FileStore))
}

// GCalAuthHandler mocks base method
func (m *MockFactory) GCalAuthHandler(oauthConfig *oauth2.Config) GCalAuthHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GCalAuthHandler", oauthConfig)
	ret0, _ := ret[0].(GCalAuthHandler)
	return ret0
}

// GCalAuthHandler indicates an expected call of GCalAuthHandler
func (mr *MockFactoryMockRecorder) GCalAuthHandler(oauthConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GCalAuthHandler", reflect.TypeOf((*MockFactory)(nil).GCalAuthHandler), oauthConfig)
}
