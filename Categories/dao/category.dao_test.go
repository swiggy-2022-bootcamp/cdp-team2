package dao

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
	mock_dynamodbiface "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/mocks/mock_dynamodbiface"
)

func getDummycat() *models.Category {
	return &models.Category{1, models.CategoryDesc{"test", "test", "test", "test", "test"}}
}

func TestGetRandomKey(t *testing.T) {
	assert.NotZero(t, getRandomKey())
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &CategoryDao{mdb}

	res := getDummycat()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().GetItem(&dynamodb.GetItemInput{
		TableName: categoriesTableName(),
		Key:       getKeyFilter(res.CategoryId),
	}).Times(1).Return(&dynamodb.GetItemOutput{Item: output}, nil)

	actual, err := catservice.GetByID(res.CategoryId)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &CategoryDao{mdb}

	res := getDummycat()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().Scan(&dynamodb.ScanInput{
		TableName: categoriesTableName(),
	}).Times(1).Return(&dynamodb.ScanOutput{Items: []map[string]*dynamodb.AttributeValue{output}}, nil)

	actual, err := catservice.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(actual))
	assert.Equal(t, *res, actual[0])
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &CategoryDao{mdb}

	res := getDummycat()
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
	catservice := &CategoryDao{mdb}

	res := getDummycat()
	output, err := dynamodbattribute.MarshalMap(res)
	assert.NoError(t, err, "Error marshalling test output")

	mdb.EXPECT().UpdateItem(gomock.Any()).Times(1).Return(&dynamodb.UpdateItemOutput{Attributes: output}, nil)

	actual, err := catservice.UpdateByID(res.CategoryId, *res)
	assert.NoError(t, err)
	assert.Equal(t, res, actual)
}

func TestDeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdb := mock_dynamodbiface.NewMockDynamoDBAPI(ctrl)
	catservice := &CategoryDao{mdb}

	mdb.EXPECT().DeleteItem(&dynamodb.DeleteItemInput{
		TableName: categoriesTableName(),
		Key:       getKeyFilter(1),
	}).Times(1).Return(nil, nil)

	err := catservice.DeleteByID(1)
	assert.NoError(t, err)
}
