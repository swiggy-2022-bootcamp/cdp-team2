package dao

import (
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/db"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/literals"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
)

const (
	customEpoch = 1300000000
	tabelName   = "team-2-orders"
)

type OrderDao struct {
	db dynamodbiface.DynamoDBAPI
}

func GetOrderDao() IDao {
	return &OrderDao{
		db.GetInstance(),
	}
}

func ordersTableName() *string {
	return aws.String(tabelName)
}

func getKeyFilter(id string) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"orderId": {
			S: aws.String(id),
		},
	}
}

func getKeyFilterStatus(status string) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"status": {
			S: aws.String(status),
		},
	}
}

/**
 * Get uuid key for orderId
 * @param id Number
 */
func getRandomKey() string {
	uuid := uuid.New()
	fmt.Println(uuid.String())
	return uuid.String()
}

/**
 * Get order details of given orderId
 * @param id String
 */
func (cd *OrderDao) GetByID(id string) (*models.Order, error) {

	res, err := cd.db.GetItem(&dynamodb.GetItemInput{
		TableName: ordersTableName(),
		Key:       getKeyFilter(id),
	})

	if err != nil {
		log.Printf("Error Fetching Order: %s", err)
		return nil, err
	}

	order := models.Order{}

	if res.Item == nil {
		return nil, errors.New(literals.OrderNotFound)
	}

	if err = dynamodbattribute.UnmarshalMap(res.Item, &order); err != nil {
		log.Printf("Error unMarshalling Order: %s", err)
		return nil, err
	}

	return &order, nil
}

/**
 * Get order details of given status
 * @param statusId Number
 */
func (cd *OrderDao) GetByStatus(status int) ([]models.Order, error) {

	filt := expression.Name("status").Equal(expression.Value(status))

	expr, err := expression.NewBuilder().
		WithFilter(filt).
		Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.ScanInput{
		TableName:                 ordersTableName(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := cd.db.Scan(input)

	if err != nil {
		return nil, err
	}

	var results []models.Order

	if len(res.Items) == 0 {
		return nil, errors.New(literals.OrderNotFound)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &results); err != nil {
		log.Printf("Error unMarshalling Order: %s", err)
		return nil, err
	}

	return results, nil
}

/**
 * Get order details of given customerId
 * @param customerId String
 */

func (cd *OrderDao) GetByCustomerId(customerId string) ([]models.Order, error) {

	filt := expression.Name("customerId").Equal(expression.Value(customerId))

	expr, err := expression.NewBuilder().
		WithFilter(filt).
		Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.ScanInput{
		TableName:                 ordersTableName(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	res, err := cd.db.Scan(input)

	if err != nil {
		return nil, err
	}

	var results []models.Order

	if len(res.Items) == 0 {
		return nil, errors.New(literals.OrderNotFound)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &results); err != nil {
		log.Printf("Error unMarshalling Order: %s", err)
		return nil, err
	}

	return results, nil
}

/**
 * Get all created orders
 * @param
 */
func (cd *OrderDao) GetAll() ([]models.Order, error) {

	var results []models.Order

	res, err := cd.db.Scan(&dynamodb.ScanInput{
		TableName: ordersTableName(),
	})

	if err != nil {
		log.Printf("Error Fetching Orders %s", err.Error())
		return nil, err
	}

	if res.Items == nil {
		return results, errors.New(literals.OrderNotFound)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &results)
	if err != nil {
		log.Printf("Failed to unmarshal Orders, %v", err)
		return nil, err
	}

	return results, nil
}

/**
 * Create a order
 * @param order models.Order
 */
func (cd *OrderDao) Create(order models.Order) (*models.Order, error) {

	newId := getRandomKey()

	log.Printf("new order %+v", order)

	var exp expression.UpdateBuilder
	exp = exp.Set(expression.Name("productDesc"), expression.Value(order.ProductDesc))
	exp = exp.Set(expression.Name("status"), expression.Value(order.Status))
	exp = exp.Set(expression.Name("customerId"), expression.Value(order.CustomerId))
	exp = exp.Set(expression.Name("totalPrice"), expression.Value(order.TotalPrice))
	exp = exp.Set(expression.Name("payedPrice"), expression.Value(order.PayedPrice))

	express, _ := expression.NewBuilder().WithUpdate(exp).Build()

	createItemIn := dynamodb.UpdateItemInput{
		TableName:                 ordersTableName(),
		Key:                       getKeyFilter(newId),
		ExpressionAttributeNames:  express.Names(),
		ExpressionAttributeValues: express.Values(),
		UpdateExpression:          express.Update(),
		ReturnValues:              aws.String("ALL_NEW"),
	}

	log.Printf("create item input : %+v", createItemIn)

	resp, err := cd.db.UpdateItem(&createItemIn)

	if err != nil {
		log.Printf("error while creating order %s", err.Error())
		return nil, err
	}

	log.Printf("Order create resp %v", resp)

	savedOrder := models.Order{}
	if err = dynamodbattribute.UnmarshalMap(resp.Attributes, &savedOrder); err != nil {
		log.Printf("error while creating order %s", err.Error())
		return nil, err
	}

	return &savedOrder, nil
}

/**
 * Update order of given orderId
 * @param customerId String
 * @param order models.Order
 */
func (cd *OrderDao) UpdateByID(id string, order models.Order) (*models.Order, error) {

	toUpd, err := getOrderUpdExp(order)

	if err != nil {
		log.Printf("error while creating expression %s", err)
		return nil, err
	}

	updItemIn := dynamodb.UpdateItemInput{
		TableName:                 ordersTableName(),
		Key:                       getKeyFilter(id),
		ExpressionAttributeNames:  toUpd.Names(),
		ExpressionAttributeValues: toUpd.Values(),
		UpdateExpression:          toUpd.Update(),
		ReturnValues:              aws.String("ALL_NEW"),
		ConditionExpression:       toUpd.Condition(),
	}

	log.Printf("update item input : %+v", updItemIn)

	resp, err := cd.db.UpdateItem(&updItemIn)

	if err != nil {
		log.Printf("error while updating orders %s", err.Error())
		return nil, err
	}

	log.Printf("Order update resp %v", resp)

	updatedAttributes := models.Order{}
	if err = dynamodbattribute.UnmarshalMap(resp.Attributes, &updatedAttributes); err != nil {
		log.Printf("error while updating orders %s", err.Error())
		return nil, err
	}

	return &updatedAttributes, nil
}

/**
 * Delete order of given orderId
 * @param id String
 */
func (cd *OrderDao) DeleteByID(id string) error {

	res, err := cd.db.GetItem(&dynamodb.GetItemInput{
		TableName: ordersTableName(),
		Key:       getKeyFilter(id),
	})

	if err != nil {
		log.Printf("Error Fetching Order: %s", err)
		return err
	}

	if res.Item == nil {
		return errors.New(literals.OrderNotFound)
	}

	resp, err := cd.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: ordersTableName(),
		Key:       getKeyFilter(id),
	})

	log.Printf("delete Order resp %+v", resp)

	if err != nil {
		log.Printf("delete Order err %s", err.Error())
		return err
	}
	return nil
}

// getOrderUpdExp creates a Expression from OrderUpdateInput
// this iterates through all fields and form update expression
func getOrderUpdExp(order models.Order) (expression.Expression, error) {
	var updateExp expression.UpdateBuilder

	desc := order.ProductDesc
	if len(desc) != 0 {
		updateExp = updateExp.Set(expression.Name("productDesc"), expression.Value(desc))
	}

	if order.Status != 0 {
		updateExp = updateExp.Set(expression.Name("status"), expression.Value(order.Status))
	}

	if order.AddressId != 0 {
		updateExp = updateExp.Set(expression.Name("addressId"), expression.Value(order.AddressId))
	}

	if order.PayedPrice != 0 {
		updateExp = updateExp.Set(expression.Name("payedPrice"), expression.Value(order.PayedPrice))
	}

	exp, err := expression.
		NewBuilder().
		WithCondition(expression.AttributeExists(expression.Name("orderId"))).
		WithUpdate(updateExp).
		Build()

	return exp, err
}
