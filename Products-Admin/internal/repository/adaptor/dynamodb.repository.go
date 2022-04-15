package adaptor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/products-admin-service/internal/core/domain"
	"github.com/products-admin-service/internal/core/ports"
	"github.com/products-admin-service/pkg/errors"
)

type ProductsRepository struct {
	DynamoDBClient *dynamodb.DynamoDB
	LogMode        bool
	TableName      string
}

// check if ProductsRepository implements IProductsRepository
var _ ports.IProductsRepository = (*ProductsRepository)(nil)

func NewProductsRepository(dynamodbClient *dynamodb.DynamoDB, tableName string) *ProductsRepository {
	return &ProductsRepository{
		DynamoDBClient: dynamodbClient,
		LogMode:        false,
		TableName:      tableName,
	}
}

func (pr *ProductsRepository) Health() bool {
	_, err := pr.DynamoDBClient.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (pr *ProductsRepository) GetAll(condition expression.Expression) (products []*domain.Product, err *errors.AppError) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(pr.TableName),
	}

	response, dErr := pr.DynamoDBClient.Scan(input)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}

	// Unmarshal dynamodb map to domain.Product
	for _, _dbProduct := range response.Items {
		var _product domain.Product
		dErr = dynamodbattribute.UnmarshalMap(_dbProduct, &_product)
		if dErr != nil {
			return nil, errors.Wrap(dErr)
		}
		products = append(products, &_product)
	}
	return products, nil
}

func (pr *ProductsRepository) AddProduct(product *domain.Product) (*domain.Product, *errors.AppError) {
	entityParsed, dErr := dynamodbattribute.MarshalMap(product)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(pr.TableName),
	}
	_, dErr = pr.DynamoDBClient.PutItem(input)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	return product, nil
}

func (pr *ProductsRepository) UpdateProduct(product *domain.Product) (*domain.Product, *errors.AppError) {
	entityParsed, dErr := dynamodbattribute.MarshalMap(product)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(pr.TableName),
	}
	_, dErr = pr.DynamoDBClient.PutItem(input)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	return product, nil
}

func (pr *ProductsRepository) DeleteProduct(condition map[string]*dynamodb.AttributeValue) (response *dynamodb.DeleteItemOutput, err *errors.AppError) {
	conditionParsed, dErr := dynamodbattribute.MarshalMap(condition)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(pr.TableName),
	}
	response, dErr = pr.DynamoDBClient.DeleteItem(input)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	return response, nil
}
