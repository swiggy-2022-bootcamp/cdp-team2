package services

import (
	"fmt"
	"net/http"
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

func (ps *ProductsServices) GetProductsByCategoryId(categoryID int64) ([]*domain.Product, *errors.AppError) {

	// Convert int64 category to string type
	// categoryIDStr := *aws.String(strconv.Itoa(int(categoryID)))

	// Create the condition builder
	// filter := expression.Contains(expression.Name("product_category"), categoryIDStr)
	filter := expression.Name("_id").NotEqual(expression.Value(""))

	// Build the condtion
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return []*domain.Product{}, errors.Wrap(err)
	}

	// Call to repository
	_products, err2 := ps.ProductsRepository.GetProductsByCondition(condition)
	if err != nil {
		return []*domain.Product{}, err2
	}

	// Service level filtering for category specific products
	var _catProducts []*domain.Product
	for _, _product := range _products {
		for _, catID := range _product.ProductCategory {
			if catID == categoryID {
				_catProducts = append(_catProducts, _product)
				break
			}
		}
	}

	// Return products
	return _catProducts, nil
}
