package services

import (
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

func (ps *ProductsServices) AddProduct(*domain.Product) (string, *errors.AppError) {
	return "", errors.New("")
}

func (ps *ProductsServices) UpdateProduct(primitive.ObjectID, *domain.Product) *errors.AppError {
	return errors.New("")
}

func (ps *ProductsServices) DeleteProduct(primitive.ObjectID) *errors.AppError {
	return errors.New("")
}

func (ps *ProductsServices) GetProducts() ([]*domain.Product, *errors.AppError) {
	return nil, errors.New("")
}
