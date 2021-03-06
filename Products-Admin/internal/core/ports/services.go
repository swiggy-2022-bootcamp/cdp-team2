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
	IsProductExists(int64) (bool, *errors.AppError)
	GetProductsByCategoryId(int64) ([]*domain.Product, *errors.AppError)
	CheckProductsWithCategory(int64) (bool, *errors.AppError)
	SearchByKeyword(string) ([]*domain.Product, *errors.AppError)
	SearchByManufacturerID(int64) ([]*domain.Product, *errors.AppError)
	CheckoutProducts([]*domain.ProductIDAndQnty) ([]*domain.ProductIDMsg, []*domain.ProductIDMsg, *errors.AppError)
	SearchByStartPrice(string) ([]*domain.Product, *errors.AppError)
}
