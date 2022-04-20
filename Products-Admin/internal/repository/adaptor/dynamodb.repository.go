package adaptor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/errors"
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

func (pr *ProductsRepository) GetProductsByCondition(condition expression.Expression) (products []*domain.Product, err *errors.AppError) {
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
	input := &dynamodb.DeleteItemInput{
		Key:       condition,
		TableName: aws.String(pr.TableName),
	}
	response, dErr := pr.DynamoDBClient.DeleteItem(input)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	return response, nil
}

func (pr *ProductsRepository) GetProductById(condition map[string]*dynamodb.AttributeValue) (product *domain.Product, err *errors.AppError) {
	input := &dynamodb.GetItemInput{
		Key:       condition,
		TableName: aws.String(pr.TableName),
	}
	response, dErr := pr.DynamoDBClient.GetItem(input)
	if dErr != nil {
		return nil, errors.Wrap(dErr)
	}
	if response.Item == nil {
		return nil, errors.New("Product not found.", 404)
	}
	_product := domain.Product{}
	// if err := dynamodbattribute.UnmarshalMap(response.Item, &_product); err != nil {
	// 	return nil, errors.New(err.Error(), http.StatusInternalServerError)
	// }
	if err := dynamodbattribute.UnmarshalMap(response.Item, &_product); err != nil {
		return nil, errors.Wrap(err)
	}
	return &_product, nil
}

func (pr *ProductsRepository) IsProductExists(condition map[string]*dynamodb.AttributeValue) (exists bool, err *errors.AppError) {
	input := &dynamodb.GetItemInput{
		Key:       condition,
		TableName: aws.String(pr.TableName),
	}
	response, dErr := pr.DynamoDBClient.GetItem(input)
	if dErr != nil {
		return false, errors.Wrap(dErr)
	}
	if response.Item == nil {
		return false, nil
	}
	return true, nil
}
