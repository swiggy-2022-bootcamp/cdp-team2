package ports

import (
	"github.com/products-admin-service/internal/core/domain"
	"github.com/products-admin-service/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductsServices interface {
	AddProduct(*domain.Product) (string, *errors.AppError)
	UpdateProduct(primitive.ObjectID, *domain.Product) *errors.AppError
	DeleteProduct(primitive.ObjectID) *errors.AppError
	GetProducts() ([]*domain.Product, *errors.AppError)
}
