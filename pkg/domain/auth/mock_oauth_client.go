// Code generated by MockGen. DO NOT EDIT.
// Source: oauth_client.go

// Package auth is a generated GoMock package.
package auth

import (
	gomock "github.com/golang/mock/gomock"
	oauth2 "golang.org/x/oauth2"
	reflect "reflect"
)

// MockOAuthClient is a mock of OAuthClient interface
type MockOAuthClient struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthClientMockRecorder
}

// MockOAuthClientMockRecorder is the mock recorder for MockOAuthClient
type MockOAuthClientMockRecorder struct {
	mock *MockOAuthClient
}

// NewMockOAuthClient creates a new mock instance
func NewMockOAuthClient(ctrl *gomock.Controller) *MockOAuthClient {
	mock := &MockOAuthClient{ctrl: ctrl}
	mock.recorder = &MockOAuthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOAuthClient) EXPECT() *MockOAuthClientMockRecorder {
	return m.recorder
}

// GetAuthCodeURL mocks base method
func (m *MockOAuthClient) GetAuthCodeURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthCodeURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAuthCodeURL indicates an expected call of GetAuthCodeURL
func (mr *MockOAuthClientMockRecorder) GetAuthCodeURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthCodeURL", reflect.TypeOf((*MockOAuthClient)(nil).GetAuthCodeURL))
}

// Exchange mocks base method
func (m *MockOAuthClient) Exchange(authCode string) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exchange", authCode)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange
func (mr *MockOAuthClientMockRecorder) Exchange(authCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockOAuthClient)(nil).Exchange), authCode)
}
