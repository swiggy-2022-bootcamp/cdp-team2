package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
)

type ProductsServices struct {
	ProductsRepository ports.IProductsRepository
}

var _ ports.IProductsServices = (*ProductsServices)(nil)

const (
	customEpoch = 1300000000
)

func getRandomKey() int64 {
	return int64(time.Now().Unix() - customEpoch)
}

func NewProductsServices(productsRepository ports.IProductsRepository) *ProductsServices {
	return &ProductsServices{
		ProductsRepository: productsRepository,
	}
}

func (ps *ProductsServices) AddProduct(product *domain.Product) (int64, *errors.AppError) {
	product.ID = getRandomKey()
	if _, err := ps.ProductsRepository.AddProduct(product); err != nil {
		return 0, err
	}
	return product.ID, nil
}

func (ps *ProductsServices) UpdateProduct(productID int64, product *domain.Product) *errors.AppError {
	product.ID = productID
	if _, err := ps.ProductsRepository.UpdateProduct(product); err != nil {
		return err
	}
	return nil
}

func (ps *ProductsServices) DeleteProduct(productID int64) *errors.AppError {
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(productID)),
		},
	}
	if _, err := ps.ProductsRepository.DeleteProduct(condition); err != nil {
		return err
	}
	return nil
}

func (ps *ProductsServices) GetProducts() ([]*domain.Product, *errors.AppError) {
	var _products []*domain.Product
	filter := expression.Name("_id").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return _products, errors.New(err.Error(), http.StatusInternalServerError)
	}
	_products, dErr := ps.ProductsRepository.GetProductsByCondition(condition)
	if dErr != nil {
		return _products, dErr
	}
	return _products, nil
}

func (ps *ProductsServices) GetProductById(productID int64) (*domain.Product, *errors.AppError) {
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(productID)),
		},
	}
	_product, err := ps.ProductsRepository.GetProductById(condition)
	if err != nil {
		return nil, err
	}
	return _product, err
}

func (ps *ProductsServices) IsProductExists(productID int64) (bool, *errors.AppError) {
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(fmt.Sprint(productID)),
		},
	}
	exists, err := ps.ProductsRepository.IsProductExists(condition)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (ps *ProductsServices) CheckProductsWithCategory(categoryID int64) (bool, *errors.AppError) {

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
	_products, err2 := ps.ProductsRepository.GetProductsByScanInput(&queryInput)
	if err2 != nil {
		return false, err2
	}

	// Check if products exists or not
	if len(_products) == 0 {
		return false, nil
	}
	return true, nil
}

func (ps *ProductsServices) GetProductsByCategoryId(categoryID int64) ([]*domain.Product, *errors.AppError) {

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
	_products, err2 := ps.ProductsRepository.GetProductsByScanInput(&queryInput)
	if err2 != nil {
		return []*domain.Product{}, err2
	}

	// Return products
	return _products, nil
}

func (ps *ProductsServices) SearchByKeyword(keyword string) ([]*domain.Product, *errors.AppError) {

	// Define the filter expression for searching product by keyword
	filter1 := expression.Contains(expression.Name("model"), keyword)
	filter2 := expression.Contains(expression.Name("model"), strings.ToUpper(keyword))
	filter3 := expression.Contains(expression.Name("sku"), keyword)
	filter4 := expression.Contains(expression.Name("sku"), strings.ToUpper(keyword))
	filter := filter1.Or(filter2).Or(filter3).Or(filter4)

	// Build expression from above filter
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return []*domain.Product{}, errors.Wrap(err)
	}

	// Get products from repository
	_products, err2 := ps.ProductsRepository.GetProductsByCondition(condition)
	if err2 != nil {
		return []*domain.Product{}, err2
	}

	// Return the products to handlers
	return _products, nil
}

func (ps *ProductsServices) SearchByManufacturerID(manufacturerID int64) ([]*domain.Product, *errors.AppError) {

	// Define filter expression
	filter := expression.Name("manufacturer_id").Equal(expression.Value(manufacturerID))

	// Build condition from above filter
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return []*domain.Product{}, errors.Wrap(err)
	}

	// Get products from dynamodb repository
	_products, err2 := ps.ProductsRepository.GetProductsByCondition(condition)
	if err2 != nil {
		return []*domain.Product{}, err2
	}

	// Return products
	return _products, nil
}
