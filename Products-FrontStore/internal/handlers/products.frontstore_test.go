package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/pkg/errors"
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

func TestGetProductsSuccess(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/", nil)

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().GetProductList().Return([]*domain.Product{}, nil)

	// initialize handlers
	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProductList(context)

	assert.Equal(t, response.Code, http.StatusOK, "Successfull Get Products should return ok status")
}

func TestGetProductsError(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/", nil)

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	// define expected behaviour
	productsMockServices.EXPECT().GetProductList().Return([]*domain.Product{}, errors.New("Something went wrong", http.StatusInternalServerError))

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProductList(context)

	assert.Equal(t, response.Code, http.StatusInternalServerError, "Failed get products should return internal server error status.")
}

func TestGetProductsByIDSuccess(t *testing.T) {

	productID := int64(1234)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/", nil)
	context.Params = gin.Params{
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
	productsMockServices.EXPECT().GetProductById(productID).Return(&_mockProduct, nil)

	// initialize handlers
	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProductById(context)

	assert.Equal(t, response.Code, http.StatusOK, "Successfull Get Product by ID should return ok status")
}

func TestGetProductByIDError(t *testing.T) {

	productID := int64(3232)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/", nil)
	context.Params = gin.Params{
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
	productsMockServices.EXPECT().GetProductById(productID).Return(nil, errors.New("Something went wrong", http.StatusInternalServerError))

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProductById(context)

	assert.Equal(t, response.Code, http.StatusInternalServerError, "Failed get product by id should return internal server error status.")
}

func TestGetProductByInvalidID(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/", nil)
	context.Params = gin.Params{
		{
			Key:   "id",
			Value: "DKfd",
		},
	}

	// setting up mock services
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

	productsHandlers := NewHandlers(productsMockServices)

	// funtion to test
	productsHandlers.GetProductById(context)

	assert.Equal(t, response.Code, http.StatusBadRequest, "Failed get product by id should return bad request error status.")
}

func TestSearchByCategoryInvalidID(t *testing.T) {

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/category/", nil)
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
	productsHandlers.GetProductListByCategoryId(context)

	assert.Equal(t, response.Code, http.StatusBadRequest, "Invalid category ID should return bad request response.")
}

func TestSearchByCategoryNotExists(t *testing.T) {

	categoryID := int64(23)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/rest/products/category", nil)
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
	productsHandlers.GetProductListByCategoryId(context)

	assert.Equal(t, response.Code, http.StatusInternalServerError, "Failed search should return internal server response.")
}

func TestSearchByCategorySuccess(t *testing.T) {

	categoryID := int64(23)

	// Test context
	gin.SetMode(gin.TestMode)
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	context.Request = httptest.NewRequest(http.MethodGet, "/api/admin/products/category", nil)
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
	productsMockServices.EXPECT().GetProductsByCategoryId(categoryID).Return([]*domain.Product{&_mockProduct}, nil)

	// funtion to test
	productsHandlers.GetProductListByCategoryId(context)

	assert.Equal(t, response.Code, http.StatusOK, "Successfull search should return ok response.")
}
