// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/providers/cloudflare_workers_kv.go

// Package mock_providers is a generated GoMock package.
package mock_providers

import (
	context "context"
	reflect "reflect"

	cloudflare "github.com/cloudflare/cloudflare-go"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudflareClient is a mock of CloudflareClient interface.
type MockCloudflareClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudflareClientMockRecorder
}

// MockCloudflareClientMockRecorder is the mock recorder for MockCloudflareClient.
type MockCloudflareClientMockRecorder struct {
	mock *MockCloudflareClient
}

// NewMockCloudflareClient creates a new mock instance.
func NewMockCloudflareClient(ctrl *gomock.Controller) *MockCloudflareClient {
	mock := &MockCloudflareClient{ctrl: ctrl}
	mock.recorder = &MockCloudflareClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudflareClient) EXPECT() *MockCloudflareClientMockRecorder {
	return m.recorder
}

// ListWorkersKVs mocks base method.
func (m *MockCloudflareClient) ListWorkersKVs(ctx context.Context, namespaceID string) (cloudflare.ListStorageKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkersKVs", ctx, namespaceID)
	ret0, _ := ret[0].(cloudflare.ListStorageKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkersKVs indicates an expected call of ListWorkersKVs.
func (mr *MockCloudflareClientMockRecorder) ListWorkersKVs(ctx, namespaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkersKVs", reflect.TypeOf((*MockCloudflareClient)(nil).ListWorkersKVs), ctx, namespaceID)
}

// ReadWorkersKV mocks base method.
func (m *MockCloudflareClient) ReadWorkersKV(ctx context.Context, namespaceID, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWorkersKV", ctx, namespaceID, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWorkersKV indicates an expected call of ReadWorkersKV.
func (mr *MockCloudflareClientMockRecorder) ReadWorkersKV(ctx, namespaceID, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWorkersKV", reflect.TypeOf((*MockCloudflareClient)(nil).ReadWorkersKV), ctx, namespaceID, key)
}

// WriteWorkersKV mocks base method.
func (m *MockCloudflareClient) WriteWorkersKV(ctx context.Context, namespaceID, key string, value []byte) (cloudflare.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteWorkersKV", ctx, namespaceID, key, value)
	ret0, _ := ret[0].(cloudflare.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteWorkersKV indicates an expected call of WriteWorkersKV.
func (mr *MockCloudflareClientMockRecorder) WriteWorkersKV(ctx, namespaceID, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteWorkersKV", reflect.TypeOf((*MockCloudflareClient)(nil).WriteWorkersKV), ctx, namespaceID, key, value)
}

// WriteWorkersKVBulk mocks base method.
func (m *MockCloudflareClient) WriteWorkersKVBulk(ctx context.Context, namespaceID string, kvs cloudflare.WorkersKVBulkWriteRequest) (cloudflare.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteWorkersKVBulk", ctx, namespaceID, kvs)
	ret0, _ := ret[0].(cloudflare.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteWorkersKVBulk indicates an expected call of WriteWorkersKVBulk.
func (mr *MockCloudflareClientMockRecorder) WriteWorkersKVBulk(ctx, namespaceID, kvs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteWorkersKVBulk", reflect.TypeOf((*MockCloudflareClient)(nil).WriteWorkersKVBulk), ctx, namespaceID, kvs)
}
