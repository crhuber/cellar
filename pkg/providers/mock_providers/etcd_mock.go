// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/providers/etcd.go

// Package mock_providers is a generated GoMock package.
package mock_providers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// MockEtcdClient is a mock of EtcdClient interface.
type MockEtcdClient struct {
	ctrl     *gomock.Controller
	recorder *MockEtcdClientMockRecorder
}

// MockEtcdClientMockRecorder is the mock recorder for MockEtcdClient.
type MockEtcdClientMockRecorder struct {
	mock *MockEtcdClient
}

// NewMockEtcdClient creates a new mock instance.
func NewMockEtcdClient(ctrl *gomock.Controller) *MockEtcdClient {
	mock := &MockEtcdClient{ctrl: ctrl}
	mock.recorder = &MockEtcdClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEtcdClient) EXPECT() *MockEtcdClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockEtcdClient) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*clientv3.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockEtcdClientMockRecorder) Get(ctx, key interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEtcdClient)(nil).Get), varargs...)
}

// Put mocks base method.
func (m *MockEtcdClient) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, val}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*clientv3.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockEtcdClientMockRecorder) Put(ctx, key, val interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, val}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockEtcdClient)(nil).Put), varargs...)
}