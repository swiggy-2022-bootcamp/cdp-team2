package dao

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"

	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/db"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/literals"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	customEpoch = 1300000000
	tabelName   = "team-2-shippingAddress"
)

type ShippingDao struct {
	db dynamodbiface.DynamoDBAPI
}

func GetShippingDao() IDao {
	return &ShippingDao{
		db.GetInstance(),
	}
}

func ordersTableName() *string {
	return aws.String(tabelName)
}

func getKeyFilter(id int) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"addressId": {
			N: aws.String(fmt.Sprint(id)),
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

func getRandomKey() int {
	return int(time.Now().Unix() - customEpoch)
}

/**
 * Get shipping address of a customer
 * @param customerId int
 */
func (cd *ShippingDao) GetByCustomerId(customerId string) ([]models.ShippingAddress, error) {

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

	var results []models.ShippingAddress

	if len(res.Items) == 0 {
		return nil, errors.New(literals.AddressNotFound)
	}

	if err = dynamodbattribute.UnmarshalListOfMaps(res.Items, &results); err != nil {
		log.Printf("Error unMarshalling Address: %s", err)
		return nil, err
	}

	return results, nil
}

/**
 * Create a shipping address of a customer
 * @parama ddress models.ShippingAddress
 */
func (cd *ShippingDao) Create(shippingAddress models.ShippingAddress) (*models.ShippingAddress, error) {

	newId := getRandomKey()

	var exp expression.UpdateBuilder
	exp = exp.Set(expression.Name("firstName"), expression.Value(shippingAddress.FirstName))
	exp = exp.Set(expression.Name("lastName"), expression.Value(shippingAddress.LastName))
	exp = exp.Set(expression.Name("city"), expression.Value(shippingAddress.City))
	exp = exp.Set(expression.Name("addressLine1"), expression.Value(shippingAddress.AddressLine1))
	exp = exp.Set(expression.Name("addressLine2"), expression.Value(shippingAddress.AddressLine2))
	exp = exp.Set(expression.Name("postCode"), expression.Value(shippingAddress.PostCode))
	exp = exp.Set(expression.Name("countryCode"), expression.Value(shippingAddress.CountryCode))
	exp = exp.Set(expression.Name("customerId"), expression.Value(shippingAddress.CustomerId))

	express, _ := expression.NewBuilder().WithUpdate(exp).Build()

	createItemIn := dynamodb.UpdateItemInput{
		TableName:                 ordersTableName(),
		Key:                       getKeyFilter(newId),
		ExpressionAttributeNames:  express.Names(),
		ExpressionAttributeValues: express.Values(),
		UpdateExpression:          express.Update(),
		ReturnValues:              aws.String("ALL_NEW"),
	}

	resp, err := cd.db.UpdateItem(&createItemIn)

	if err != nil {
		log.Printf("error while creating shipping-address %s", err.Error())
		return nil, err
	}

	savedAddress := models.ShippingAddress{}
	if err = dynamodbattribute.UnmarshalMap(resp.Attributes, &savedAddress); err != nil {
		log.Printf("error while creating shipping-address %s", err.Error())
		return nil, err
	}

	return &savedAddress, nil
}
