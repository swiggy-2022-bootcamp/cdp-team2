package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

	_resProduct = domain.Product{
		ID:             int64(120),
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

func TestAddProduct(t *testing.T) {

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().AddProduct(&_mockProduct).Return(&_mockProduct, nil)

	// Call test function
	res, err := mockProductsServices.AddProduct(&_mockProduct)

	assert.Nil(t, err, "Successfull add product should not return any error.")
	assert.NotNil(t, res, "Successfull add product should return product ID.")

}

func TestAddProductError(t *testing.T) {

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().AddProduct(&_mockProduct).Return(nil, errors.New("Something went wrong", http.StatusInternalServerError))

	// Call test function
	res, err := mockProductsServices.AddProduct(&_mockProduct)

	assert.NotNil(t, err, "Failed add product should not return error.")
	assert.Equal(t, res, int64(0), "Failed add product should return 0 product ID.")

}

func TestUpdateProduct(t *testing.T) {

	dummyID := int64(2349372)

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().UpdateProduct(&_mockProduct).Return(&_mockProduct, nil)

	// Call test function
	err := mockProductsServices.UpdateProduct(dummyID, &_mockProduct)

	assert.Nil(t, err, "Successfull update product should not return any error.")

}

func TestUpdateProductError(t *testing.T) {

	dummyID := int64(2349372)

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().UpdateProduct(&_mockProduct).Return(nil, errors.New("Product Not found", http.StatusNotFound))

	// Call test function
	err := mockProductsServices.UpdateProduct(dummyID, &_mockProduct)

	assert.NotNil(t, err, "Failed update product should return error.")

}

func TestDeleteProduct(t *testing.T) {

	dummyID := int64(2349372)

	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(dummyID)),
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().DeleteProduct(condition).Return(nil, nil)

	// Call test function
	err := mockProductsServices.DeleteProduct(dummyID)

	assert.Nil(t, err, "Successfull delete product should not return any error.")

}

func TestDeleteProductError(t *testing.T) {

	dummyID := int64(2349372)

	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(dummyID)),
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().DeleteProduct(condition).Return(nil, errors.New("Product Not found", http.StatusNotFound))

	// Call test function
	err := mockProductsServices.DeleteProduct(dummyID)

	assert.NotNil(t, err, "Failed delete product should return error.")

}

func TestGetProducts(t *testing.T) {

	// Dummy Filter condition
	filter := expression.Name("_id").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByCondition(condition).Return([]*domain.Product{}, nil)

	// Call test function
	products, err := mockProductsServices.GetProducts()

	assert.Nil(t, err, "Successfull get products should not return any error.")
	assert.NotNil(t, products, "Successfull get products should return products.")
}

func TestGetProductsError(t *testing.T) {

	// Dummy Filter condition
	filter := expression.Name("_id").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByCondition(condition).Return(nil, errors.New("Something went wrong", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.GetProducts()

	assert.NotNil(t, err, "Failed get products should return any error.")
	assert.Nil(t, products, "Failed get products should not return products.")
}

func TestGetProductByID(t *testing.T) {

	dummyID := int64(2349372)

	// Dummy Condition
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(dummyID)),
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductById(condition).Return(&_mockProduct, nil)

	// Call test function
	product, err := mockProductsServices.GetProductById(dummyID)

	assert.Nil(t, err, "Successfull get product should not return any error.")
	assert.NotNil(t, product, "Successfull get product should return product.")
}

func TestGetProductByIDError(t *testing.T) {

	dummyID := int64(2349372)
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(dummyID)),
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductById(condition).Return(nil, errors.New("Product Not found", http.StatusNotFound))

	// Call test function
	product, err := mockProductsServices.GetProductById(dummyID)

	assert.NotNil(t, err, "Failed get product should return error.")
	assert.Nil(t, product, "Failed get product should not return product.")

}

func TestIsProductExists(t *testing.T) {

	dummyID := int64(2349372)

	// Dummy Condition
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(dummyID)),
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().IsProductExists(condition).Return(true, nil)

	// Call test function
	exists, err := mockProductsServices.IsProductExists(dummyID)

	assert.Nil(t, err, "Successfull response should not return any error.")
	assert.Equal(t, exists, true, "Successfull response should return true.")
}

func TestIsProductExistsError(t *testing.T) {

	dummyID := int64(2349372)
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(dummyID)),
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().IsProductExists(condition).Return(false, errors.New("Product Not found", http.StatusNotFound))

	// Call test function
	exists, err := mockProductsServices.IsProductExists(dummyID)

	assert.NotNil(t, err, "Failed response should return error.")
	assert.Equal(t, exists, false, "Failed response should return false.")

}

func TestCheckProductsByCategoryID(t *testing.T) {

	categoryID := int64(12432)

	// Create queryInput
	queryInput := dynamodb.ScanInput{
		TableName:        nil,
		FilterExpression: aws.String("contains (product_category, :categoryID)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":categoryID": {
				N: aws.String(strconv.Itoa(int(categoryID))),
			},
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByScanInput(&queryInput).Return([]*domain.Product{&_mockProduct}, nil)

	// Call test function
	isExists, err := mockProductsServices.CheckProductsWithCategory(categoryID)

	assert.Nil(t, err, "Successfull  products search should not return any error.")
	assert.Equal(t, isExists, true, "Successfull products search should return true.")
}

func TestCheckProductsByCatgoryIDError(t *testing.T) {

	categoryID := int64(12432)

	// Create queryInput
	queryInput := dynamodb.ScanInput{
		TableName:        nil,
		FilterExpression: aws.String("contains (product_category, :categoryID)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":categoryID": {
				N: aws.String(strconv.Itoa(int(categoryID))),
			},
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByScanInput(&queryInput).Return(nil, errors.New("Something went wrong.", http.StatusInternalServerError))

	// Call test function
	isExists, err := mockProductsServices.CheckProductsWithCategory(categoryID)

	assert.NotNil(t, err, "Failed products search should return any error.")
	assert.Equal(t, isExists, false, "Failed products search should not return true.")
}

func TestSearchProductsByCategoryID(t *testing.T) {

	categoryID := int64(12432)

	// Create queryInput
	queryInput := dynamodb.ScanInput{
		TableName:        nil,
		FilterExpression: aws.String("contains (product_category, :categoryID)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":categoryID": {
				N: aws.String(strconv.Itoa(int(categoryID))),
			},
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByScanInput(&queryInput).Return([]*domain.Product{&_mockProduct}, nil)

	// Call test function
	products, err := mockProductsServices.GetProductsByCategoryId(categoryID)

	assert.Nil(t, err, "Successfull products search should not return any error.")
	assert.NotNil(t, products, " Successfull products search should return products.")
}

func TestSearchProductsByCatgoryIDError(t *testing.T) {

	categoryID := int64(12432)

	// Create queryInput
	queryInput := dynamodb.ScanInput{
		TableName:        nil,
		FilterExpression: aws.String("contains (product_category, :categoryID)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":categoryID": {
				N: aws.String(strconv.Itoa(int(categoryID))),
			},
		},
	}

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByScanInput(&queryInput).Return(nil, errors.New("Something went wrong.", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.GetProductsByCategoryId(categoryID)

	assert.NotNil(t, err, "Failed products search should return any error.")
	assert.Equal(t, len(products), 0, "Failed products search should not return products.")
}

func TestSearchProductsByKeyword(t *testing.T) {

	keyword := "dummy"

	// Define the filter expression for searching product by keyword
	filter1 := expression.Contains(expression.Name("model"), keyword)
	filter2 := expression.Contains(expression.Name("model"), strings.ToUpper(keyword))
	filter3 := expression.Contains(expression.Name("sku"), keyword)
	filter4 := expression.Contains(expression.Name("sku"), strings.ToUpper(keyword))
	filter := filter1.Or(filter2).Or(filter3).Or(filter4)

	// Build expression from above filter
	condition, _ := expression.NewBuilder().WithFilter(filter).Build()

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByCondition(condition).Return([]*domain.Product{&_mockProduct}, nil)

	// Call test function
	products, err := mockProductsServices.SearchByKeyword(keyword)

	assert.Nil(t, err, "Successfull products search should not return any error.")
	assert.NotNil(t, products, " Successfull products search should return products.")
}

func TestSearchProductsByKeywordError(t *testing.T) {

	keyword := "dummy"

	// Define the filter expression for searching product by keyword
	filter1 := expression.Contains(expression.Name("model"), keyword)
	filter2 := expression.Contains(expression.Name("model"), strings.ToUpper(keyword))
	filter3 := expression.Contains(expression.Name("sku"), keyword)
	filter4 := expression.Contains(expression.Name("sku"), strings.ToUpper(keyword))
	filter := filter1.Or(filter2).Or(filter3).Or(filter4)

	// Build expression from above filter
	condition, _ := expression.NewBuilder().WithFilter(filter).Build()

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByCondition(condition).Return(nil, errors.New("Something went wrong.", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.SearchByKeyword(keyword)

	assert.NotNil(t, err, "Failed products search should return any error.")
	assert.Equal(t, len(products), 0, "Failed products search should not return products.")
}

func TestSearchProductsByManufacturerID(t *testing.T) {

	// Define filter expression
	manufacturerID := int64(1234)
	filter := expression.Name("manufacturer_id").Equal(expression.Value(manufacturerID))

	// Build condition from above filter
	condition, _ := expression.NewBuilder().WithFilter(filter).Build()

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByCondition(condition).Return([]*domain.Product{&_mockProduct}, nil)

	// Call test function
	products, err := mockProductsServices.SearchByManufacturerID(manufacturerID)

	assert.Nil(t, err, "Successfull products search should not return any error.")
	assert.NotNil(t, products, " Successfull products search should return products.")
}

func TestSearchProductsByManufacturerIDError(t *testing.T) {

	// Define filter expression
	manufacturerID := int64(1234)
	filter := expression.Name("manufacturer_id").Equal(expression.Value(manufacturerID))

	// Build condition from above filter
	condition, _ := expression.NewBuilder().WithFilter(filter).Build()

	// Gomock controller
	productsMockCtrl := gomock.NewController(t)
	defer productsMockCtrl.Finish()

	// Initialize mock product reposity
	mockProductRepository := mocks.NewMockIProductsRepository(productsMockCtrl)

	// Initialize mock product services
	mockProductsServices := NewProductsServices(mockProductRepository)

	// Define expected behavior of product repository
	mockProductRepository.EXPECT().GetProductsByCondition(condition).Return(nil, errors.New("Something went wrong.", http.StatusInternalServerError))

	// Call test function
	products, err := mockProductsServices.SearchByManufacturerID(manufacturerID)

	assert.NotNil(t, err, "Failed products search should return any error.")
	assert.Equal(t, len(products), 0, "Failed products search should not return products.")
}
