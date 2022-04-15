package ports

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/products-admin-service/internal/core/domain"
	"github.com/products-admin-service/pkg/errors"
)

type IProductsRepository interface {
	Health() bool
	GetAll(condition expression.Expression) (products []*domain.Product, err *errors.AppError)
	AddProduct(product *domain.Product) (*domain.Product, *errors.AppError)
	UpdateProduct(product *domain.Product) (*domain.Product, *errors.AppError)
	DeleteProduct(condition map[string]interface{}) (response *dynamodb.DeleteItemOutput, err *errors.AppError)
}
