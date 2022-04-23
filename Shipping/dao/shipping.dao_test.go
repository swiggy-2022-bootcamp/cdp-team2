package dao

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/mocks/mock_dynamodbiface"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	//mock_dynamodbiface "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/mocks/mock_dynamodbiface"
)

func getDummyShipping() *models.ShippingAddress {
	return &models.ShippingAddress{1, 2, "Taranjeet", "Singh", "Varanasi", "Zamania", "Ghazipur", "IN", 232331}
}

func TestGetRandomKey(t *testing.T) {
	assert.NotZero(t, getRandomKey())
}

//func TestGetByCustomerId(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
//	catservice := &ShippingDao{mdb}
//
//	res := []models.ShippingAddress{*getDummyShipping()}
//	output, err := dynamodbattribute.MarshalMap(res)
//	assert.NoError(t, err, "Error marshalling test output")
//
//	filt := expression.Name("customerId").Equal(expression.Value(res[0].CustomerId))
//
//	expr, err := expression.NewBuilder().
//		WithFilter(filt).
//		Build()
//
//	mdb.EXPECT().GetItem(&dynamodb.ScanInput{
//		TableName:                 ordersTableName(),
//		ExpressionAttributeNames:  expr.Names(),
//		ExpressionAttributeValues: expr.Values(),
//		FilterExpression:          expr.Filter(),
//	}).Times(1).Return(&dynamodb.ScanOutput{Items: []map[string]*dynamodb.AttributeValue{output}}, nil)
//
//	actual, err := catservice.GetByCustomerId(res[0].CustomerId)
//
//	assert.NoError(t, err)
//	assert.Equal(t, res, actual)
//}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &ShippingDao{mdb}

	res := getDummyShipping()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().UpdateItem(gomock.Any()).Times(1).Return(&dynamodb.UpdateItemOutput{Attributes: output}, nil)

	actual, err := catservice.Create(*res)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}
