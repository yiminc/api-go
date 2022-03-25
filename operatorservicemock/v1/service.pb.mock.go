// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: operatorservice/v1/service.pb.go

// Package operatorservicemock is a generated GoMock package.
package operatorservicemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	operatorservice "go.temporal.io/api/operatorservice/v1"
	grpc "google.golang.org/grpc"
)

// MockOperatorServiceClient is a mock of OperatorServiceClient interface.
type MockOperatorServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorServiceClientMockRecorder
}

// MockOperatorServiceClientMockRecorder is the mock recorder for MockOperatorServiceClient.
type MockOperatorServiceClientMockRecorder struct {
	mock *MockOperatorServiceClient
}

// NewMockOperatorServiceClient creates a new mock instance.
func NewMockOperatorServiceClient(ctrl *gomock.Controller) *MockOperatorServiceClient {
	mock := &MockOperatorServiceClient{ctrl: ctrl}
	mock.recorder = &MockOperatorServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorServiceClient) EXPECT() *MockOperatorServiceClientMockRecorder {
	return m.recorder
}

// AddSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) AddSearchAttributes(ctx context.Context, in *operatorservice.AddSearchAttributesRequest, opts ...grpc.CallOption) (*operatorservice.AddSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.AddSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSearchAttributes indicates an expected call of AddSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) AddSearchAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).AddSearchAttributes), varargs...)
}

// DeleteNamespace mocks base method.
func (m *MockOperatorServiceClient) DeleteNamespace(ctx context.Context, in *operatorservice.DeleteNamespaceRequest, opts ...grpc.CallOption) (*operatorservice.DeleteNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNamespace", varargs...)
	ret0, _ := ret[0].(*operatorservice.DeleteNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockOperatorServiceClientMockRecorder) DeleteNamespace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockOperatorServiceClient)(nil).DeleteNamespace), varargs...)
}

// ListSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) ListSearchAttributes(ctx context.Context, in *operatorservice.ListSearchAttributesRequest, opts ...grpc.CallOption) (*operatorservice.ListSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSearchAttributes indicates an expected call of ListSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) ListSearchAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListSearchAttributes), varargs...)
}

// RemoveSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) RemoveSearchAttributes(ctx context.Context, in *operatorservice.RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*operatorservice.RemoveSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.RemoveSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSearchAttributes indicates an expected call of RemoveSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) RemoveSearchAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).RemoveSearchAttributes), varargs...)
}

// MockOperatorServiceServer is a mock of OperatorServiceServer interface.
type MockOperatorServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorServiceServerMockRecorder
}

// MockOperatorServiceServerMockRecorder is the mock recorder for MockOperatorServiceServer.
type MockOperatorServiceServerMockRecorder struct {
	mock *MockOperatorServiceServer
}

// NewMockOperatorServiceServer creates a new mock instance.
func NewMockOperatorServiceServer(ctrl *gomock.Controller) *MockOperatorServiceServer {
	mock := &MockOperatorServiceServer{ctrl: ctrl}
	mock.recorder = &MockOperatorServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorServiceServer) EXPECT() *MockOperatorServiceServerMockRecorder {
	return m.recorder
}

// AddSearchAttributes mocks base method.
func (m *MockOperatorServiceServer) AddSearchAttributes(arg0 context.Context, arg1 *operatorservice.AddSearchAttributesRequest) (*operatorservice.AddSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSearchAttributes", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.AddSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSearchAttributes indicates an expected call of AddSearchAttributes.
func (mr *MockOperatorServiceServerMockRecorder) AddSearchAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSearchAttributes", reflect.TypeOf((*MockOperatorServiceServer)(nil).AddSearchAttributes), arg0, arg1)
}

// DeleteNamespace mocks base method.
func (m *MockOperatorServiceServer) DeleteNamespace(arg0 context.Context, arg1 *operatorservice.DeleteNamespaceRequest) (*operatorservice.DeleteNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.DeleteNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockOperatorServiceServerMockRecorder) DeleteNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockOperatorServiceServer)(nil).DeleteNamespace), arg0, arg1)
}

// ListSearchAttributes mocks base method.
func (m *MockOperatorServiceServer) ListSearchAttributes(arg0 context.Context, arg1 *operatorservice.ListSearchAttributesRequest) (*operatorservice.ListSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSearchAttributes", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.ListSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSearchAttributes indicates an expected call of ListSearchAttributes.
func (mr *MockOperatorServiceServerMockRecorder) ListSearchAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSearchAttributes", reflect.TypeOf((*MockOperatorServiceServer)(nil).ListSearchAttributes), arg0, arg1)
}

// RemoveSearchAttributes mocks base method.
func (m *MockOperatorServiceServer) RemoveSearchAttributes(arg0 context.Context, arg1 *operatorservice.RemoveSearchAttributesRequest) (*operatorservice.RemoveSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSearchAttributes", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.RemoveSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSearchAttributes indicates an expected call of RemoveSearchAttributes.
func (mr *MockOperatorServiceServerMockRecorder) RemoveSearchAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSearchAttributes", reflect.TypeOf((*MockOperatorServiceServer)(nil).RemoveSearchAttributes), arg0, arg1)
}
