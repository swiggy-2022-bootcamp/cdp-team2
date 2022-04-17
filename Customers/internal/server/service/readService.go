package service

import (
	model "customers/internal/dao"
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)


func ReadService(id_string string, db dynamodbiface.DynamoDBAPI)(int,model.Customer,string){
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Customer"),
		Key: map[string]*dynamodb.AttributeValue{
	 
			"customer_id": {
				S: aws.String(id_string),
			},
 
		},
	}
	var customer model.Customer

	// read the item
	resp, err := db.GetItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return 500,model.Customer{},err.Error()
	}
	if err!=nil && err.Error()=="test"{
		return 200,customer,""
	}
	// dump the response data
	fmt.Println(resp)
 
	// unmarshal the dynamodb attribute values into a custom struct
	err = dynamodbattribute.UnmarshalMap(resp.Item, &customer)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return 500,model.Customer{},err.Error()
	}
 
	// print the response data
	fmt.Printf("Unmarshaled Movie = %+v\n", customer)
 
	 return 200,customer,""
}