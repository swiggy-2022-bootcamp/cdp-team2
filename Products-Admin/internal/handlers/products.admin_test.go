package handlers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
// 	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/mocks"
// )

// func TestAddProduct(t *testing.T) {

// 	// Mock product
// 	_product := &domain.Product{}
// 	_productStr, _ := json.Marshal(_product)

// 	// Test context
// 	gin.SetMode(gin.TestMode)
// 	response := httptest.NewRecorder()
// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = httptest.NewRequest(http.MethodPost, "/api/rest_admin/products", bytes.NewBuffer(_productStr))

// 	// setting up mock services
// 	productsMockCtrl := gomock.NewController(t)
// 	defer productsMockCtrl.Finish()

// 	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

// 	// define expected behaviour
// 	productsMockServices.EXPECT().AddProduct(_product).Return(0, nil)

// 	productsHandlers := NewHandlers(productsMockServices)

// 	// funtion to test
// 	productsHandlers.AddProduct(context)

// 	if response.Code != http.StatusCreated {
// 		t.Error("Successfull add product  should return 201 response.")
// 	}
// }

// func TestUpdateProduct(t *testing.T) {

// 	// Mock product
// 	_product := &domain.Product{}
// 	_productStr, _ := json.Marshal(_product)

// 	// Test context
// 	gin.SetMode(gin.TestMode)
// 	response := httptest.NewRecorder()
// 	context, _ := gin.CreateTestContext(response)
// 	context.Request = httptest.NewRequest(http.MethodPost, "/api/rest_admin/products/", bytes.NewBuffer(_productStr))

// 	// setting up mock services
// 	productsMockCtrl := gomock.NewController(t)
// 	defer productsMockCtrl.Finish()

// 	productsMockServices := mocks.NewMockIProductsServices(productsMockCtrl)

// 	// define expected behaviour
// 	productsMockServices.EXPECT().AddProduct(_product).Return(0, nil)

// 	productsHandlers := NewHandlers(productsMockServices)

// 	// funtion to test
// 	productsHandlers.AddProduct(context)

// 	if response.Code != http.StatusCreated {
// 		t.Error("Successfull add product  should return 201 response.")
// 	}
// }
