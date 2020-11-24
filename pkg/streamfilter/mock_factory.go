// Code generated by MockGen. DO NOT EDIT.
// Source: ./factory.go

// Package streamfilter is a generated GoMock package.
package streamfilter

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "mosn.io/api"
	reflect "reflect"
)

// MockStreamFilterFactory is a mock of StreamFilterFactory interface.
type MockStreamFilterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockStreamFilterFactoryMockRecorder
}

// MockStreamFilterFactoryMockRecorder is the mock recorder for MockStreamFilterFactory.
type MockStreamFilterFactoryMockRecorder struct {
	mock *MockStreamFilterFactory
}

// NewMockStreamFilterFactory creates a new mock instance.
func NewMockStreamFilterFactory(ctrl *gomock.Controller) *MockStreamFilterFactory {
	mock := &MockStreamFilterFactory{ctrl: ctrl}
	mock.recorder = &MockStreamFilterFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamFilterFactory) EXPECT() *MockStreamFilterFactoryMockRecorder {
	return m.recorder
}

// CreateFilterChain mocks base method.
func (m *MockStreamFilterFactory) CreateFilterChain(context context.Context, callbacks api.StreamFilterChainFactoryCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateFilterChain", context, callbacks)
}

// CreateFilterChain indicates an expected call of CreateFilterChain.
func (mr *MockStreamFilterFactoryMockRecorder) CreateFilterChain(context, callbacks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilterChain", reflect.TypeOf((*MockStreamFilterFactory)(nil).CreateFilterChain), context, callbacks)
}

// GetConfig mocks base method.
func (m *MockStreamFilterFactory) GetConfig() StreamFiltersConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(StreamFiltersConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockStreamFilterFactoryMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockStreamFilterFactory)(nil).GetConfig))
}