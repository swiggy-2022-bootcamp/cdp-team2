package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
)

var (
	_mockProduct = domain.Product{
		Model:          "dummy model name",
		Quantity:       10,
		Price:          "1000",
		Minimum:        1,
		ManufacturerID: 20,
		Sku:            "dummy sku",
		ProductSeoUrl: &domain.ProductSeoUrl{
			Keyword:    "dummy keyword",
			StoreID:    20,
			LanguageID: 2,
		},
		Points:             300,
		Rewards:            4000,
		Image:              "dummy image url",
		ShippingID:         20,
		Weight:             40,
		Width:              20,
		Length:             10,
		Height:             5,
		ProductRelated:     []int64{20, 10},
		ProductCategory:    []int64{14, 20, 40},
		ProductDescription: []*domain.ProductDescription{},
	}
)

func TestAddProductEmtpyBody(t *testing.T) {

	// Mock empty product
	_product := &domain.Product{}
	_productStr, _ := json.Marshal(_product)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/api/rest_admin/products", bytes.NewBuffer(_productStr))

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.AddProduct(context)

	if response.Code != http.StatusBadRequest {
		t.Error("Empty body should generate validation error and return bad request status.")
	}
}

func TestAddProductSuccess(t *testing.T) {

	// Mock product
	_productStr, _ := json.Marshal(_mockProduct)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPost, "/api/rest_admin/products/", bytes.NewBuffer(_productStr))

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().AddProduct(&_mockProduct).Return(int64(2392), nil)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.AddProduct(context)

	if response.Code != http.StatusCreated {
		t.Error("Successfull add product should return 201 response.")
	}
}

func TestUpdateProductEmtpyBody(t *testing.T) {

	// Mock empty product
	_product := &domain.Product{}
	_productStr, _ := json.Marshal(_product)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPut, "/api/rest_admin/products", bytes.NewBuffer(_productStr))

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.UpdateProduct(context)

	if response.Code != http.StatusBadRequest {
		t.Error("Empty body should generate validation error and return bad request status.")
	}
}

func TestUpdateProductSuccess(t *testing.T) {

	// Mock product
	_productStr, _ := json.Marshal(_mockProduct)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodPut, "/api/rest_admin/products/", bytes.NewBuffer(_productStr))
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "2392",
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().UpdateProduct(int64(2392), &_mockProduct).Return(nil)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.UpdateProduct(context)

	if response.Code != http.StatusCreated {
		t.Error("Successfull update product should return 201 response.")
	}
}

func TestDeleteProductNotExists(t *testing.T) {

	// mock product id
	productID := int64(2392)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodDelete, "/api/rest_admin/products/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprint(productID),
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().DeleteProduct(productID).Return(errors.New("Product not exists", http.StatusInternalServerError))

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.DeleteProduct(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("Non exists product should return internal server error response.")
	}
}

func TestDeleteProductInvalidID(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodDelete, "/api/rest_admin/products/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "",
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.DeleteProduct(context)

	if response.Code != http.StatusBadRequest {
		t.Error("None product ID should return bad request response.")
	}
}

func TestDeleteProductSuccess(t *testing.T) {

	// mock product id
	productID := int64(2392)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodDelete, "/api/rest_admin/products/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprint(productID),
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().DeleteProduct(productID).Return(nil)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.DeleteProduct(context)

	if response.Code != http.StatusOK {
		t.Error("Successfull delete product should return ok status.")
	}
}

func TestGetProductsSuccess(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/", nil)

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().GetProducts().Return([]*domain.Product{}, nil)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProducts(context)

	if response.Code != http.StatusOK {
		t.Error("Successfull get products should return ok status.")
	}
}

func TestGetProductsError(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/", nil)

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().GetProducts().Return([]*domain.Product{}, errors.New("Something went wrong", http.StatusInternalServerError))

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProducts(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("Failed get products should return internal server error status.")
	}
}

func TestSearchByCategoryInvalidID(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/seach/category/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "iaminvalid",
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.SearchByCategoryID(context)

	if response.Code != http.StatusBadRequest {
		t.Error("Invalid category ID should return bad request response.")
	}
}

func TestSearchByCategoryNotExists(t *testing.T) {

	categoryID := int64(23)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprint(categoryID),
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// Define expected behavior
	productsMockServices.EXPECT().GetProductsByCategoryId(categoryID).Return(nil, errors.New("Something went wrong.", http.StatusInternalServerError))

	// funtion to test
	productsHandlers.SearchByCategoryID(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("Failed search should return internal server response.")
	}
}

func TestSearchByCategorySuccess(t *testing.T) {

	categoryID := int64(23)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/seach/manufacturer", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprint(categoryID),
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// Define expected behavior
	productsMockServices.EXPECT().GetProductsByCategoryId(categoryID).Return([]*domain.Product{}, nil)

	// funtion to test
	productsHandlers.SearchByCategoryID(context)

	if response.Code != http.StatusOK {
		t.Error("Successfull search should return ok response.")
	}
}

func TestSearchByManufacturerInvalidID(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/search/manufacturer/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: "iaminvalid",
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.SearchByManufacturerID(context)

	if response.Code != http.StatusBadRequest {
		t.Error("Invalid manufacturer ID should return bad request response.")
	}
}

func TestSearchByManufacturerNotExists(t *testing.T) {

	manufacturerID := int64(233)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/search/manufacturer/", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprint(manufacturerID),
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// Define expected behavior
	productsMockServices.EXPECT().SearchByManufacturerID(manufacturerID).Return(nil, errors.New("Something went wrong.", http.StatusInternalServerError))

	// funtion to test
	productsHandlers.SearchByManufacturerID(context)

	if response.Code != http.StatusInternalServerError {
		t.Error("Failed search should return internal server response.")
	}
}

func TestSearchByManufacturerSuccess(t *testing.T) {

	manufacturerID := int64(23)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/search/manufacturer", nil)
	context.Params = []gin.Param{
		{
			Key:   "id",
			Value: fmt.Sprint(manufacturerID),
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// Define expected behavior
	productsMockServices.EXPECT().SearchByManufacturerID(manufacturerID).Return([]*domain.Product{}, nil)

	// funtion to test
	productsHandlers.SearchByManufacturerID(context)

	if response.Code != http.StatusOK {
		t.Error("Successfull search should return ok response.")
	}
}

func TestSearchByEmptyKeyword(t *testing.T) {

	keyword := ""

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/search/keyword/", nil)
	context.Params = []gin.Param{
		{
			Key:   "keyword",
			Value: "",
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// define expected behavior
	productsMockServices.EXPECT().SearchByKeyword(keyword).Return(nil, errors.New("Please provide keyword.", http.StatusBadRequest))

	// funtion to test
	productsHandlers.SearchByKeyword(context)

	if response.Code != http.StatusBadRequest {
		t.Error("Invalid keyword should return bad request response.")
	}
}

func TestSearchByKeyword(t *testing.T) {

	keyword := "device"

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest_admin/products/search/keyword/", nil)
	context.Params = []gin.Param{
		{
			Key:   "keyword",
			Value: keyword,
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// define expected behavior
	productsMockServices.EXPECT().SearchByKeyword(keyword).Return([]*domain.Product{}, nil)

	// funtion to test
	productsHandlers.SearchByKeyword(context)

	if response.Code != http.StatusOK {
		t.Error("Successfull keyword should return ok response.")
	}
}
