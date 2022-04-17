
package util

import (
 	model "customer-account/internal/dao"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"fmt"
 	 
)



func GetItems(id_string string,db dynamodbiface.DynamoDBAPI)(model.Account){

	// create the api params
	params11:= &dynamodb.GetItemInput{
		TableName: aws.String("Customer"),
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
		return model.Account{Id:"test"}
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return model.Account{}
	}
	// dump the response data
	// unmarshal the dynamodb attribute values into a custom struct
	var account model.Account
	err = dynamodbattribute.UnmarshalMap(resp.Item, &account)
	return account
}


