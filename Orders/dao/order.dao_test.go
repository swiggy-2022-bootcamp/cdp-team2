package dao

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	mock_dynamodbiface "github.com/swiggy-2022-bootcamp/cdp-team2/Order/mocks/mock_dynamodbiface"
)

func getDummyOrder() *models.Order {
	productDesc := []models.ProductDesc{}
	productDesc = append(productDesc, models.ProductDesc{1, 5, 5, 10, 200})
	return &models.Order{"1", "2", 2, 2, 200, 190, productDesc}
}

func TestGetRandomKey(t *testing.T) {
	assert.NotZero(t, getRandomKey())
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	res := getDummyOrder()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().GetItem(&dynamodb.GetItemInput{
		TableName: ordersTableName(),
		Key:       getKeyFilter(res.OrderId),
	}).Times(1).Return(&dynamodb.GetItemOutput{Item: output}, nil)

	actual, err := catservice.GetByID(res.OrderId)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestGetByCustomerId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	res := []models.Order{*getDummyOrder()}
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	filt := expression.Name("customerId").Equal(expression.Value(res[0].CustomerId))

	expr, err := expression.NewBuilder().
		WithFilter(filt).
		Build()

	mdb.EXPECT().GetItem(&dynamodb.ScanInput{
		TableName:                 ordersTableName(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}).Times(1).Return(&dynamodb.GetItemOutput{Item: output}, nil)

	actual, err := catservice.GetByCustomerId(res[0].CustomerId)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestGetByStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	res := []models.Order{*getDummyOrder()}
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	filt := expression.Name("status").Equal(expression.Value(res[0].Status))

	expr, err := expression.NewBuilder().
		WithFilter(filt).
		Build()

	mdb.EXPECT().GetItem(&dynamodb.ScanInput{
		TableName:                 ordersTableName(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}).Times(1).Return(&dynamodb.GetItemOutput{Item: output}, nil)

	actual, err := catservice.GetByStatus(int(res[0].Status))
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	res := getDummyOrder()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().Scan(&dynamodb.ScanInput{
		TableName: ordersTableName(),
	}).Times(1).Return(&dynamodb.ScanOutput{Items: []map[string]*dynamodb.AttributeValue{output}}, nil)

	actual, err := catservice.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, *res, actual[0])
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	res := getDummyOrder()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().UpdateItem(gomock.Any()).Times(1).Return(&dynamodb.UpdateItemOutput{Attributes: output}, nil)

	actual, err := catservice.Create(*res)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestUpdateByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	res := getDummyOrder()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().UpdateItem(gomock.Any()).Times(1).Return(&dynamodb.UpdateItemOutput{Attributes: output}, nil)

	actual, err := catservice.UpdateByID(res.OrderId, *res)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestDeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &OrderDao{mdb}

	mdb.EXPECT().DeleteItem(&dynamodb.DeleteItemInput{
		TableName: ordersTableName(),
		Key:       getKeyFilter("1"),
	}).Times(1).Return(nil, nil)

	err := catservice.DeleteByID("1")
	assert.NoError(t, err)
}
