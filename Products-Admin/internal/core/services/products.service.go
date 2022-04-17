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
	_products, dErr := ps.ProductsRepository.GetAll(condition)
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
