package dao_test

import (
	"github.com/golang/mock/gomock"
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	mock "customers/mock"
	"customers/dao/models"
	"customers/dao"
	// "fmt"
	"strconv"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"errors"
	// "github.com/stretchr/testify/assert"
)

// var mockDynamo=mock_dynamodbiface.MockDynamoDBAPIMockRecorder{}
// var dynamoService=mock_dynamodbiface.NewMockDynamoDBAPI(&mockDynamo)
// func GetMockCustomerDao() dao.IDao{
// 	return &daCustomerDao{
// 		mockDynamo,
// 	}
// }
const (tableName="Customer")
func RandomAddress()models.Address{
	return models.Address{
		Firstname:"uday",
		Lastname:"lastname",
		AddressLine1: "House no: 5-6-49",
		AddressLine2: "Vidhya nagar colony",
		City:"Kamareddy",
		PostCode:503111,
		CountryCode:"+91",
	}
}
func RandomCustomer()models.Customer{
	return models.Customer{
		Id              :"test",
		Firstname       :"Uday",
		Lastname        :"Kiran",
		Email           :"test@gmail.com",
		Password        :"password",
		Confirm         :"password",
		Telephone       :"123-234-43",
		Fax				: "123-234-43",
		Status			: "1",
		Approved		: "1",
		Safe			: "1",
		CustomerGroupID : 1,
		DateAdded       :    "test",
		Address			: []models.Address{RandomAddress()},
	}
}
 


func TestGet(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)
	customer:=RandomCustomer()


	params := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customer.Id),
			},
		},
	}
	dynamodbattribute.MarshalMap(customer)
	mockDynamo.EXPECT().GetItem(params).Return(&dynamodb.GetItemOutput{},nil).AnyTimes()

	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err:=cd.Get(customer.Id)
	assert.Equal(t,err,errors.New("Customer Not Found"))
}


func TestGetByEmail(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)
	customer:=RandomCustomer()


	params := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
        IndexName: aws.String("email-index"),
        KeyConditions: map[string]*dynamodb.Condition{
            "email": {
                ComparisonOperator: aws.String("EQ"),
                AttributeValueList: []*dynamodb.AttributeValue{
                    {
                        S: aws.String(customer.Email),
                    },
                },
            },
        },

	}
	dynamodbattribute.MarshalMap(customer)
	mockDynamo.EXPECT().Query(params).Return(&dynamodb.QueryOutput{},nil).AnyTimes()

	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err:=cd.GetByEmail(customer.Email)
	assert.Equal(t,err,errors.New("Customer with this email does not exist!!"))
}

func TestDelete(t *testing.T){
	customer:=RandomCustomer()
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customer.Id),
			},
		 
		},
	}
	mockDynamo.EXPECT().DeleteItem(params).Return(&dynamodb.DeleteItemOutput{},nil)
	cd:=&dao.CustomerDao{mockDynamo}
	_,err:=cd.Delete(customer.Id)
	assert.Nil(t,err)
}

func TestCreate(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)


	customer:=RandomCustomer()
 
	params := &dynamodb.QueryInput{
		TableName: aws.String("Customer"),
        IndexName: aws.String("email-index"),
        KeyConditions: map[string]*dynamodb.Condition{
            "email": {
                ComparisonOperator: aws.String("EQ"),
                AttributeValueList: []*dynamodb.AttributeValue{
                    {
                        S: aws.String(customer.Email),
                    },
                },
            },
        },

	}

	info, err := dynamodbattribute.MarshalMap(customer)
	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(tableName),
	}
	mockDynamo.EXPECT().Query(params).Return(&dynamodb.QueryOutput{},nil)
	mockDynamo.EXPECT().PutItem(input).Return(&dynamodb.PutItemOutput{},nil)

	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err=cd.Create(customer)
	assert.Nil(t,err)
}



func TestUpdate(t *testing.T){
	mockCtrl:=gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDynamo:=mock.NewMockDynamoDBAPI(mockCtrl)
	customer:=RandomCustomer()
	customer.Email=""
	obj:=customer
	addresslist,_:=dynamodbattribute.MarshalList(obj.Address)
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customer.Id),
			}, 
		},
		UpdateExpression: aws.String("set firstname=:f, lastname=:l, password=:p, confirm=:c, telephone=:t, fax=:fax, newsletter=:nl , approved=:ap, safe=:sf, customer_group_id=:cgi, address=:adrs, statuss=:stas "),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {S: aws.String(obj.Firstname)},
			":l": {S: aws.String(obj.Lastname)},
			":p": {S:aws.String(obj.Password)},
			":c": {S:aws.String(obj.Confirm)},
			":t": {S:aws.String(obj.Telephone)},
			":fax":{S:aws.String(obj.Fax)},
			":nl":{S:aws.String(obj.Newsletter)},
			":ap":{S:aws.String(obj.Approved)},
			":sf":{S:aws.String(obj.Safe)},
			":cgi":{N:aws.String(strconv.Itoa(obj.CustomerGroupID))},
			":adrs":{L:addresslist},
			":stas":{S:aws.String(obj.Status)},
 		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}

	params2:= &dynamodb.GetItemInput{
		TableName: aws.String("Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(obj.Id),
			},

		},
	}
	
	mockDynamo.EXPECT().GetItem(params2).Return(&dynamodb.GetItemOutput{},nil).AnyTimes()

	mockDynamo.EXPECT().UpdateItem(params).Return(&dynamodb.UpdateItemOutput{},nil)
	var cd dao.IDao
	cd=&dao.CustomerDao{mockDynamo}
	_,err:=cd.Update(customer.Id,customer)
	assert.Nil(t,err)
	
}


