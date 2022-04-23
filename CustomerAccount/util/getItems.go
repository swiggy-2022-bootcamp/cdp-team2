
package util

import (
 	"customer-account/dao/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"fmt"
 	 
)



func GetItems(id_string string,db dynamodbiface.DynamoDBAPI)(models.Account){

	// create the api params
	params11:= &dynamodb.GetItemInput{
		TableName: aws.String("team-2-Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			// "_id": {
			// 	N: aws.String(strconv.Itoa(12345)),
			// },
			"customer_id": {
				S: aws.String(id_string),
			},

		},
	}

	// read the item
	resp, err := db.GetItem(params11)

	if err!=nil && err.Error()=="test"{
		return models.Account{Id:"test"}
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return models.Account{}
	}
	// dump the response data
	// unmarshal the dynamodb attribute values into a custom struct
	var account models.Account
	err = dynamodbattribute.UnmarshalMap(resp.Item, &account)
	return account
}


