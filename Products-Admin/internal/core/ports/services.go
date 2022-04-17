package ports

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
)

type IProductsServices interface {
	AddProduct(*domain.Product) (int, *errors.AppError)
	UpdateProduct(int, *domain.Product) *errors.AppError
	DeleteProduct(int) *errors.AppError
	GetProducts() ([]*domain.Product, *errors.AppError)
}
