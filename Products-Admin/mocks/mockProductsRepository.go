// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports (interfaces: IProductsRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	expression "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	errors "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
)

// MockIProductsRepository is a mock of IProductsRepository interface.
type MockIProductsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProductsRepositoryMockRecorder
}

// MockIProductsRepositoryMockRecorder is the mock recorder for MockIProductsRepository.
type MockIProductsRepositoryMockRecorder struct {
	mock *MockIProductsRepository
}

// NewMockIProductsRepository creates a new mock instance.
func NewMockIProductsRepository(ctrl *gomock.Controller) *MockIProductsRepository {
	mock := &MockIProductsRepository{ctrl: ctrl}
	mock.recorder = &MockIProductsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductsRepository) EXPECT() *MockIProductsRepositoryMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockIProductsRepository) AddProduct(arg0 *domain.Product) (*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", arg0)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockIProductsRepositoryMockRecorder) AddProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockIProductsRepository)(nil).AddProduct), arg0)
}

// DeleteProduct mocks base method.
func (m *MockIProductsRepository) DeleteProduct(arg0 map[string]*dynamodb.AttributeValue) (*dynamodb.DeleteItemOutput, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0)
	ret0, _ := ret[0].(*dynamodb.DeleteItemOutput)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockIProductsRepositoryMockRecorder) DeleteProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockIProductsRepository)(nil).DeleteProduct), arg0)
}

// GetProductById mocks base method.
func (m *MockIProductsRepository) GetProductById(arg0 map[string]*dynamodb.AttributeValue) (*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockIProductsRepositoryMockRecorder) GetProductById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockIProductsRepository)(nil).GetProductById), arg0)
}

// GetProductsByCondition mocks base method.
func (m *MockIProductsRepository) GetProductsByCondition(arg0 expression.Expression) ([]*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByCondition", arg0)
	ret0, _ := ret[0].([]*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetProductsByCondition indicates an expected call of GetProductsByCondition.
func (mr *MockIProductsRepositoryMockRecorder) GetProductsByCondition(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByCondition", reflect.TypeOf((*MockIProductsRepository)(nil).GetProductsByCondition), arg0)
}

// GetProductsByScanInput mocks base method.
func (m *MockIProductsRepository) GetProductsByScanInput(arg0 *dynamodb.ScanInput) ([]*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByScanInput", arg0)
	ret0, _ := ret[0].([]*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// GetProductsByScanInput indicates an expected call of GetProductsByScanInput.
func (mr *MockIProductsRepositoryMockRecorder) GetProductsByScanInput(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByScanInput", reflect.TypeOf((*MockIProductsRepository)(nil).GetProductsByScanInput), arg0)
}

// Health mocks base method.
func (m *MockIProductsRepository) Health() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Health indicates an expected call of Health.
func (mr *MockIProductsRepositoryMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockIProductsRepository)(nil).Health))
}

// IsProductExists mocks base method.
func (m *MockIProductsRepository) IsProductExists(arg0 map[string]*dynamodb.AttributeValue) (bool, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProductExists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// IsProductExists indicates an expected call of IsProductExists.
func (mr *MockIProductsRepositoryMockRecorder) IsProductExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProductExists", reflect.TypeOf((*MockIProductsRepository)(nil).IsProductExists), arg0)
}

// UpdateProduct mocks base method.
func (m *MockIProductsRepository) UpdateProduct(arg0 *domain.Product) (*domain.Product, *errors.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0)
	ret0, _ := ret[0].(*domain.Product)
	ret1, _ := ret[1].(*errors.AppError)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockIProductsRepositoryMockRecorder) UpdateProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockIProductsRepository)(nil).UpdateProduct), arg0)
}
