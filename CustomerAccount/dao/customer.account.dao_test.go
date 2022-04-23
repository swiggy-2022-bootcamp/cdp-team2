package dao_test

import (
	"github.com/golang/mock/gomock"
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	mock "customer-account/mock"
	"customer-account/dao/models"
	"customer-account/dao"
	// "fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"customer-account/config"
	// "errors"
	// "github.com/stretchr/testify/assert"
)

var (
	tableName = config.AWS["TABLE_NAME"]
)

func RandomAccount()models.Account{
	return models.Account{
		Id              :"test",
		Firstname       :"Uday",
		Lastname        :"Kiran",
		Email           :"test@gmail.com",
		Password        :"password",
		Confirm         :"password",
		Telephone       :"123-234-43",
		UserBalance     : 0.00,
		CustomerGroupID : 1,
		DateAdded       :    "test",
		Agree           : "1",
	}
}

func RandomAccount2()models.Account2{
	return models.Account2{
		Id              :"test",
		Firstname       :"Uday",
		Lastname        :"Kiran",
		Email           :"test@gmail.com",
		Cart 			: models.Cart{},
		Reward			: models.Reward{},
		UserBalance     : 0.0,
		CustomerGroupID : 1,
		DateAdded       :    "test",
	}
}


func TestGet(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)
	account:=RandomAccount2()


	params := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(account.Id),
			},
		},
	}
	dynamodbattribute.MarshalMap(account)
	mockDynamo.EXPECT().GetItem(params).Return(&dynamodb.GetItemOutput{},nil).AnyTimes()

	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err:=cd.Get(account.Id)
	assert.Nil(t,err)
}



func TestCreate(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)


	account:=RandomAccount()
	dynamodbattribute.MarshalMap(account)

	params := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
        IndexName: aws.String("email-index"),
        KeyConditions: map[string]*dynamodb.Condition{
            "email": {
                ComparisonOperator: aws.String("EQ"),
                AttributeValueList: []*dynamodb.AttributeValue{
                    {
                        S: aws.String(account.Email),
                    },
                },
            },
        },

	}

	info, err := dynamodbattribute.MarshalMap(account)
	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(tableName),
	}
	mockDynamo.EXPECT().Query(params).Return(&dynamodb.QueryOutput{},nil)
	mockDynamo.EXPECT().PutItem(input).Return(&dynamodb.PutItemOutput{},nil)

	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err=cd.Create(account)
	assert.Nil(t,err)
}



func TestUpdate(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)
	account:=RandomAccount()
	account.Email=""
	obj:=account
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(account.Id),
			}, 
		},
		UpdateExpression: aws.String("set firstname=:f, lastname=:l, password=:p, confirm=:c, telephone=:t, email=:e"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {S: aws.String(obj.Firstname)},
			":l": {S: aws.String(obj.Lastname)},
			":p": {S:aws.String(obj.Password)},
			":c": {S:aws.String(obj.Confirm)},
			":t": {S:aws.String(obj.Telephone)},
			":e": {S:aws.String(obj.Email)},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}

	params2:= &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			// "_id": {
			// 	N: aws.String(strconv.Itoa(12345)),
			// },
			"customer_id": {
				S: aws.String(obj.Id),
			},

		},
	}
	
	mockDynamo.EXPECT().GetItem(params2).Return(&dynamodb.GetItemOutput{},nil).AnyTimes()

	mockDynamo.EXPECT().UpdateItem(params).Return(&dynamodb.UpdateItemOutput{},nil)
	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err:=cd.Update(account.Id,account)
	assert.Nil(t,err)
	
}


