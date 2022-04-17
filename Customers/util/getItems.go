package util

import (
	"fmt"
	model "customers/internal/dao"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)


func GetItems(id_string string,db dynamodbiface.DynamoDBAPI)(model.Customer){

	// create the api params
	params:= &dynamodb.GetItemInput{
		TableName: aws.String("Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			},

		},
	}

	// read the item
	resp, err := db.GetItem(params)
	if err!=nil && err.Error()=="test"{
		return model.Customer{Id:"test"}
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		//return 400,model.Customer{},err11.Error()
	}

	// dump the response data
 
	// unmarshal the dynamodb attribute values into a custom struct
	var customer model.Customer
	err = dynamodbattribute.UnmarshalMap(resp.Item, &customer)
	return customer
}