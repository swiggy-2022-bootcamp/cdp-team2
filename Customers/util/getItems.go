package util

import (
	"fmt"
	model "customers/dao/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)


func GetItems(id_string string,db dynamodbiface.DynamoDBAPI)(model.Customer){

	// create the api params
	params:= &dynamodb.GetItemInput{
		TableName: aws.String("team-2-Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			},

		},
	}

	// read the item
	resp, err := db.GetItem(params)
 	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return model.Customer{}
	}

	// dump the response data
 
	// unmarshal the dynamodb attribute values into a custom struct
	var customer model.Customer
	err = dynamodbattribute.UnmarshalMap(resp.Item, &customer)
	return customer
}