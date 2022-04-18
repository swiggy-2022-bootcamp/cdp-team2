package ports

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
)

type IProductsRepository interface {
	Health() bool
	GetAll(condition expression.Expression) (products []*domain.Product, err *errors.AppError)
	GetProductById(map[string]*dynamodb.AttributeValue) (*domain.Product, *errors.AppError)
	IsProductExists(map[string]*dynamodb.AttributeValue) (bool, *errors.AppError)
	AddProduct(product *domain.Product) (*domain.Product, *errors.AppError)
	UpdateProduct(product *domain.Product) (*domain.Product, *errors.AppError)
	DeleteProduct(condition map[string]*dynamodb.AttributeValue) (response *dynamodb.DeleteItemOutput, err *errors.AppError)
}
