package service

import (
	model "customers/internal/dao"
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)


func GetService(email string, db dynamodbiface.DynamoDBAPI)(int,model.Customer,string){
	params := &dynamodb.QueryInput{
		TableName: aws.String("Customer"),
        IndexName: aws.String("email-index"),
        KeyConditions: map[string]*dynamodb.Condition{
            "email": {
                ComparisonOperator: aws.String("EQ"),
                AttributeValueList: []*dynamodb.AttributeValue{
                    {
                        S: aws.String(email),
                    },
                },
            },
        },

	}

	// read the item
	resp, err := db.Query(params)
	if err!=nil && err.Error()=="test"{
		return 200,model.Customer{},""
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		//return
	}

	// dump the response data
	fmt.Println(resp)

	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var account []model.Customer
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &account)
	if len(account)==0{
		return 404,model.Customer{},"No Customer with that email id";
	}


	return 200,account[0],""
}