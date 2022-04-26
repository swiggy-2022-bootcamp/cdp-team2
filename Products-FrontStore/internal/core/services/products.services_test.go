package services

import (
	"context"
	"net/http"
	"testing"

	pb "common/protos/products"

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

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductsGrpcClientSvc := mocks.NewMockIProductsGrpcClientServices(productsMockCtrl)

	// Initialize mock product services
	ctx := context.TODO()
	mockProductsServices := NewProductsServices(context.TODO(), mockProductsGrpcClientSvc)

	// Define expected behavior of product repository
	mockProductsGrpcClientSvc.EXPECT().GetAvailableProducts(ctx, &pb.EmptyRequest{}).Return(&pb.ProductsResponse{
		Products: []*pb.Product{},
	}, nil)

	// Call test function
	_, err := mockProductsServices.GetProductList()

	assert.Nil(t, err, "Successfull list products should not return any error.")
}

func TestGetProductsError(t *testing.T) {

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductsGrpcClientSvc := mocks.NewMockIProductsGrpcClientServices(productsMockCtrl)

	// Initialize mock product services
	ctx := context.TODO()
	mockProductsServices := NewProductsServices(context.TODO(), mockProductsGrpcClientSvc)

	// Define expected behavior of product repository
	mockProductsGrpcClientSvc.EXPECT().GetAvailableProducts(ctx, &pb.EmptyRequest{}).Return(nil, errors.New("Something went wrong", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.GetProductList()

	assert.NotNil(t, err, "Failed get products should return any error.")
	assert.Nil(t, products, "Failed get products should not return products.")
}

func TestGetProductByIDSuccess(t *testing.T) {

	productID := int64(13242)

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductsGrpcClientSvc := mocks.NewMockIProductsGrpcClientServices(productsMockCtrl)

	// Initialize mock product services
	ctx := context.TODO()
	mockProductsServices := NewProductsServices(context.TODO(), mockProductsGrpcClientSvc)

	// Define expected behavior of product repository
	mockProductsGrpcClientSvc.EXPECT().GetProductById(ctx, &pb.ProductIDRequest{
		ProductID: productID,
	}).Return(&pb.ProductResponse{
		Product: &pb.Product{},
	}, nil)

	// Call test function
	_, err := mockProductsServices.GetProductById(productID)

	assert.Nil(t, err, "Successfull Get product by id should not return any error.")
}

func TestGetProductByIDError(t *testing.T) {

	productID := int64(1234)

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductsGrpcClientSvc := mocks.NewMockIProductsGrpcClientServices(productsMockCtrl)

	// Initialize mock product services
	ctx := context.TODO()
	mockProductsServices := NewProductsServices(context.TODO(), mockProductsGrpcClientSvc)

	// Define expected behavior of product repository
	mockProductsGrpcClientSvc.EXPECT().GetProductById(ctx, &pb.ProductIDRequest{
		ProductID: productID,
	}).Return(nil, errors.New("Something went wrong", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.GetProductById(productID)

	assert.NotNil(t, err, "Failed get product should return any error.")
	assert.Nil(t, products, "Failed get product should not return products.")
}

func TestGetProductByCategoryIDSuccess(t *testing.T) {

	catID := int64(13242)

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductsGrpcClientSvc := mocks.NewMockIProductsGrpcClientServices(productsMockCtrl)

	// Initialize mock product services
	ctx := context.TODO()
	mockProductsServices := NewProductsServices(context.TODO(), mockProductsGrpcClientSvc)

	// Define expected behavior of product grpc client
	mockProductsGrpcClientSvc.EXPECT().GetProductsByCategoryId(ctx, &pb.CategoryIDRequest{
		CategoryID: catID,
	}).Return(&pb.ProductsResponse{
		Products: []*pb.Product{},
	}, nil)

	// Call test function
	_, err := mockProductsServices.GetProductsByCategoryId(catID)

	assert.Nil(t, err, "Successfull Get products by category id should not return any error.")
}

func TestGetProductByCategoryIDError(t *testing.T) {

	catID := int64(242)

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductsGrpcClientSvc := mocks.NewMockIProductsGrpcClientServices(productsMockCtrl)

	// Initialize mock product services
	ctx := context.TODO()
	mockProductsServices := NewProductsServices(context.TODO(), mockProductsGrpcClientSvc)

	// Define expected behavior of product repository
	mockProductsGrpcClientSvc.EXPECT().GetProductsByCategoryId(ctx, &pb.CategoryIDRequest{
		CategoryID: catID,
	}).Return(nil, errors.New("Something went wrong", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.GetProductsByCategoryId(catID)

	assert.NotNil(t, err, "Failed get products should return any error.")
	assert.Nil(t, products, "Failed get products should not return products.")
}
