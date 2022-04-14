package ports

import (
	"github.com/products-admin-service/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductsRepository interface {
	AddProduct(*domain.Product) (string, error)
	UpdateProduct(*domain.Product) error
	DeleteProduct(primitive.ObjectID) error
}
