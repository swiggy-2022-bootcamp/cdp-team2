package dao

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
)

type DynamoDAO interface {
	AddCartItem(product models.Product) *errors.ServerError
	GetCart(customerId string) (models.Cart, *errors.ServerError)
}

type dynamoDAO struct {
	client *dynamodb.DynamoDB
	config *config.WebServerConfig
}

var dynamoDao DynamoDAO
var dynamoDaoOnce sync.Once

func InitDynamoDAO(client *dynamodb.DynamoDB, config *config.WebServerConfig) DynamoDAO {
	dynamoDaoOnce.Do(func() {
		dynamoDao = &dynamoDAO{
			client: client,
			config: config,
		}
	})

	return dynamoDao
}

func GetDynamoDAO() DynamoDAO {
	if dynamoDao == nil {
		panic("Dynamo DAO not initialised")
	}

	return dynamoDao
}

// AddProduct add product details to the cart of the user
func (dao *dynamoDAO) AddCartItem(product models.Product) *errors.ServerError {
	return nil
}

func (dao *dynamoDAO) GetCart(customerId string) (models.Cart, *errors.ServerError) {
	cart := models.Cart{}

	result, err := dao.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("cart"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(customerId),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		//msg := "Could not find '" + customerId + "'"
		return cart, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &cart)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println(cart)

	return cart, nil
}
