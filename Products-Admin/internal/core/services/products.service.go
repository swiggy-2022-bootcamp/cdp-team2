package services

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/products-admin-service/internal/core/domain"
	"github.com/products-admin-service/internal/core/ports"
	"github.com/products-admin-service/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductsServices struct {
	ProductsRepository ports.IProductsRepository
}

var _ ports.IProductsServices = (*ProductsServices)(nil)

func NewProductsServices(productsRepository ports.IProductsRepository) *ProductsServices {
	return &ProductsServices{
		ProductsRepository: productsRepository,
	}
}

func (ps *ProductsServices) AddProduct(product *domain.Product) (string, *errors.AppError) {
	product.ID = primitive.NewObjectID()
	if _, err := ps.ProductsRepository.AddProduct(product); err != nil {
		return "", err
	}
	return product.ID.Hex(), nil
}

func (ps *ProductsServices) UpdateProduct(productID primitive.ObjectID, product *domain.Product) *errors.AppError {
	product.ID = productID
	if _, err := ps.ProductsRepository.UpdateProduct(product); err != nil {
		return err
	}
	return nil
}

func (ps *ProductsServices) DeleteProduct(productID primitive.ObjectID) *errors.AppError {
	condition := map[string]*dynamodb.AttributeValue{
		"_id": {
			N: aws.String(productID.Hex()),
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
