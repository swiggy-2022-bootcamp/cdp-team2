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
	AddCartItem(customerId string, product models.Product) *errors.ServerError
	GetCart(customerId string) (models.Cart, *errors.ServerError)
	DeleteCartItem(customerId string, productId string) *errors.ServerError
	EmptyCart(customerId string) *errors.ServerError
	UpdateCartItem(customerId string, productId string, newQuantity int32) *errors.ServerError
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
func (dao *dynamoDAO) AddCartItem(customerId string, product models.Product) *errors.ServerError {
	cart, err := dao.GetCart("133")
	if err != nil {
		log.Info("cart not found for customer")
		cart.CustomerId = customerId
	}

	cart.Products = append(cart.Products, product)

	err = dao.addCartDetails(cart)
	if err != nil {
		log.Error("an error occurred while adding the cart item")
		return err
	}

	log.Info("successfully added the cart item to the customer")
	return nil
}

// DeleteCartItem deletes the cart item with given productId for the given customer using customerId
func (dao *dynamoDAO) DeleteCartItem(customerId string, productId string) *errors.ServerError {
	cart, err := dao.GetCart(customerId)
	if err != nil {
		log.Error("cart not found for customer")
		return &errors.CartEmptyError
	}

	products := cart.Products
	deleteItemIdx, err := dao.getItemIndex(products, productId)
	if err != nil {
		log.Error("product with given product ID: ", productId, "does not exist in cart")
		return nil
	}

	var updatedProducts []models.Product
	for idx, product := range products {
		if idx == deleteItemIdx {
			continue
		} else {
			updatedProducts = append(updatedProducts, product)
		}
	}

	cart.Products = updatedProducts

	err = dao.addCartDetails(cart)
	if err != nil {
		log.Error("an error occurred while deleting the cart item")
		return err
	}

	log.Info("successfully deleted the cart item from cart for customer ", customerId)
	return nil
}

// UpdateCartItem updates the cart item quantity using productId for the given customer
func (dao *dynamoDAO) UpdateCartItem(customerId string, productId string, newQuantity int32) *errors.ServerError {
	cart, err := dao.GetCart(customerId)
	if err != nil {
		log.Error("cart not found for customer")
		return &errors.CartEmptyError
	}

	updateItemIdx, err := dao.getItemIndex(cart.Products, productId)
	if err != nil {
		log.Error("product with given product ID: ", productId, "does not exist in cart")
		return nil
	}

	cart.Products[updateItemIdx].Quantity = newQuantity
	err = dao.addCartDetails(cart)
	if err != nil {
		log.Error("an error occurred while updating the cart item count for customer: ", customerId)
		return err
	}

	return nil
}

// EmptyCart deletes the cart for the customer using customerId
func (dao *dynamoDAO) EmptyCart(customerId string) *errors.ServerError {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
		TableName: aws.String("cart"),
	}

	_, err := dao.client.DeleteItem(input)
	if err != nil {
		log.WithError(err).Error("an error occurred while deleting the cart for the customer: ", customerId)
		return &errors.EmptyCartFailedError
	}

	log.Info("successfully delete the cart for the customer: ", customerId)
	return nil
}

// GetCart fetches all the cart details for the given customer id
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

	log.Info("successfully got the cart for the customer: ", customerId)
	return cart, nil
}

func (dao *dynamoDAO) getItemIndex(products []models.Product, productId string) (int, *errors.ServerError) {
	for idx, product := range products {
		if product.ProductId == productId {
			return idx, nil
		}
	}

	return -1, &errors.ProductNotFoundInCart
}

func (dao *dynamoDAO) addCartDetails(cart models.Cart) *errors.ServerError {
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

	return nil
}
