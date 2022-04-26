// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/ports (interfaces: IProductsGrpcClientServices)

// Package mocks is a generated GoMock package.
package mocks

import (
	products "common/protos/products"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockIProductsGrpcClientServices is a mock of IProductsGrpcClientServices interface.
type MockIProductsGrpcClientServices struct {
	ctrl     *gomock.Controller
	recorder *MockIProductsGrpcClientServicesMockRecorder
}

// MockIProductsGrpcClientServicesMockRecorder is the mock recorder for MockIProductsGrpcClientServices.
type MockIProductsGrpcClientServicesMockRecorder struct {
	mock *MockIProductsGrpcClientServices
}

// NewMockIProductsGrpcClientServices creates a new mock instance.
func NewMockIProductsGrpcClientServices(ctrl *gomock.Controller) *MockIProductsGrpcClientServices {
	mock := &MockIProductsGrpcClientServices{ctrl: ctrl}
	mock.recorder = &MockIProductsGrpcClientServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductsGrpcClientServices) EXPECT() *MockIProductsGrpcClientServicesMockRecorder {
	return m.recorder
}

// CheckProductsWithCategory mocks base method.
func (m *MockIProductsGrpcClientServices) CheckProductsWithCategory(arg0 context.Context, arg1 *products.CategoryIDRequest, arg2 ...grpc.CallOption) (*products.BoolResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckProductsWithCategory", varargs...)
	ret0, _ := ret[0].(*products.BoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckProductsWithCategory indicates an expected call of CheckProductsWithCategory.
func (mr *MockIProductsGrpcClientServicesMockRecorder) CheckProductsWithCategory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckProductsWithCategory", reflect.TypeOf((*MockIProductsGrpcClientServices)(nil).CheckProductsWithCategory), varargs...)
}

// DeductProductsQuantity mocks base method.
func (m *MockIProductsGrpcClientServices) DeductProductsQuantity(arg0 context.Context, arg1 *products.CheckoutRequest, arg2 ...grpc.CallOption) (*products.CheckoutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeductProductsQuantity", varargs...)
	ret0, _ := ret[0].(*products.CheckoutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeductProductsQuantity indicates an expected call of DeductProductsQuantity.
func (mr *MockIProductsGrpcClientServicesMockRecorder) DeductProductsQuantity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeductProductsQuantity", reflect.TypeOf((*MockIProductsGrpcClientServices)(nil).DeductProductsQuantity), varargs...)
}

// GetAvailableProducts mocks base method.
func (m *MockIProductsGrpcClientServices) GetAvailableProducts(arg0 context.Context, arg1 *products.EmptyRequest, arg2 ...grpc.CallOption) (*products.ProductsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAvailableProducts", varargs...)
	ret0, _ := ret[0].(*products.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableProducts indicates an expected call of GetAvailableProducts.
func (mr *MockIProductsGrpcClientServicesMockRecorder) GetAvailableProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableProducts", reflect.TypeOf((*MockIProductsGrpcClientServices)(nil).GetAvailableProducts), varargs...)
}

// GetProductById mocks base method.
func (m *MockIProductsGrpcClientServices) GetProductById(arg0 context.Context, arg1 *products.ProductIDRequest, arg2 ...grpc.CallOption) (*products.ProductResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProductById", varargs...)
	ret0, _ := ret[0].(*products.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockIProductsGrpcClientServicesMockRecorder) GetProductById(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockIProductsGrpcClientServices)(nil).GetProductById), varargs...)
}

// GetProductsByCategoryId mocks base method.
func (m *MockIProductsGrpcClientServices) GetProductsByCategoryId(arg0 context.Context, arg1 *products.CategoryIDRequest, arg2 ...grpc.CallOption) (*products.ProductsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProductsByCategoryId", varargs...)
	ret0, _ := ret[0].(*products.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByCategoryId indicates an expected call of GetProductsByCategoryId.
func (mr *MockIProductsGrpcClientServicesMockRecorder) GetProductsByCategoryId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByCategoryId", reflect.TypeOf((*MockIProductsGrpcClientServices)(nil).GetProductsByCategoryId), varargs...)
}

// IsProductExists mocks base method.
func (m *MockIProductsGrpcClientServices) IsProductExists(arg0 context.Context, arg1 *products.ProductIDRequest, arg2 ...grpc.CallOption) (*products.BoolResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsProductExists", varargs...)
	ret0, _ := ret[0].(*products.BoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsProductExists indicates an expected call of IsProductExists.
func (mr *MockIProductsGrpcClientServicesMockRecorder) IsProductExists(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProductExists", reflect.TypeOf((*MockIProductsGrpcClientServices)(nil).IsProductExists), varargs...)
}