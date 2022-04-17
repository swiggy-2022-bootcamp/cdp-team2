package dao

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"

	log "github.com/sirupsen/logrus"
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
	cart, err := dao.GetCart("133")
	if err != nil {
		log.Info("cart not found for customer")
		cart.CustomerId = "133"
	}

	cart.Products = append(cart.Products, product)

	av, goErr := dynamodbattribute.MarshalMap(cart)
	if goErr != nil {
		log.WithError(goErr).Error("an error occurred while marshalling reward to dynamo attribute")
		return &errors.MarshalToAttributesError
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("cart"),
	}

	_, goErr = dao.client.PutItem(input)
	if goErr != nil {
		log.WithField("Error: ", goErr).Error("got error while calling GetItem API of dynamo db")
		return &errors.DatabaseQueryFailed
	}

	log.Info("successfully added the cart item to the customer")
	return nil
}

func (dao *dynamoDAO) GetCart(customerId string) (models.Cart, *errors.ServerError) {
	cart := models.Cart{}

	result, err := dao.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("cart"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
	})

	if err != nil {
		log.WithField("Error: ", err).Error("got error while calling GetItem API of dynamo db")
		return cart, &errors.DatabaseQueryFailed
	}

	if result.Item == nil {
		errMsg := "Could not find cart for customer: '" + customerId + "'"
		log.WithField("Error: ", errMsg).Error("record not found")
		return cart, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &cart)
	if err != nil {
		log.WithError(err).Error("an error occurred while unmarshalling dynamo attributes to reward struct")
		return cart, &errors.UnmarshalAttributesError
	}

	fmt.Println(cart)

	log.Info("Successfully got the cart for the customer: ", customerId)
	return cart, nil
}
