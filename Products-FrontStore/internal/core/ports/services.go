package ports

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/pkg/errors"
)

type IProductsServices interface {
	GetProductList() ([]*domain.Product, *errors.AppError)
	GetProductById(int64) (*domain.Product, *errors.AppError)
}
