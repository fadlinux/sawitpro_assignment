// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateEstate mocks base method.
func (m *MockServiceInterface) CreateEstate(ctx context.Context, payload CreateEstateRequest) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEstate", ctx, payload)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEstate indicates an expected call of CreateEstate.
func (mr *MockServiceInterfaceMockRecorder) CreateEstate(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEstate", reflect.TypeOf((*MockServiceInterface)(nil).CreateEstate), ctx, payload)
}

// CreateTree mocks base method.
func (m *MockServiceInterface) CreateTree(ctx context.Context, payload CreateTreeRequest) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", ctx, payload)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockServiceInterfaceMockRecorder) CreateTree(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockServiceInterface)(nil).CreateTree), ctx, payload)
}