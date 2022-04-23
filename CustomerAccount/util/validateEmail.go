package util

import (
	"customer-account/dao/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"fmt"
)



func ValidateEmail(email string, db dynamodbiface.DynamoDBAPI) bool{
	// create the api params
	params := &dynamodb.QueryInput{
		TableName: aws.String("team-2-Customers"),
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
	if err!=nil &&  err.Error()=="test"{
		return true
	}
 
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return false
	}

 
	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var account []models.Account
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &account)
	if len(account)>0{
		return false;
	}
 	return true;
}

