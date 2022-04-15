package ports

import (
	"github.com/products-admin-service/internal/core/domain"
	"github.com/products-admin-service/pkg/errors"
)

type IProductsServices interface {
	AddProduct(*domain.Product) (int, *errors.AppError)
	UpdateProduct(int, *domain.Product) *errors.AppError
	DeleteProduct(int) *errors.AppError
	GetProducts() ([]*domain.Product, *errors.AppError)
}
