package ports

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
)

type IProductsServices interface {
	AddProduct(*domain.Product) (int64, *errors.AppError)
	UpdateProduct(int64, *domain.Product) *errors.AppError
	DeleteProduct(int64) *errors.AppError
	GetProducts() ([]*domain.Product, *errors.AppError)
	GetProductById(int64) (*domain.Product, *errors.AppError)
}
