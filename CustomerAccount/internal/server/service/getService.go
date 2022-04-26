package service

import (
	model "customer-account/internal/dao"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	grpc "customer-account/internal/grpc"

)

func ReadService(id_string string,db dynamodbiface.DynamoDBAPI)(int, model.Account2,string){
		// create the api params

		fmt.Println("*********",id_string)
		params := &dynamodb.GetItemInput{
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
		resp, err := db.GetItem(params)
		if err!=nil && err.Error()=="test"{
			return 200,model.Account2{},err.Error()
		}
	
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
			return 500,model.Account2{},err.Error()
		}
	 
		// dump the response data
		fmt.Println(resp)
	 
		// unmarshal the dynamodb attribute values into a custom struct
		var account model.Account2
		err = dynamodbattribute.UnmarshalMap(resp.Item, &account)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
			return 500,model.Account2{},err.Error()
		}
	 
		// print the response data
		fmt.Printf("Unmarshaled Movie-- = %+v\n", account)
	 
		account.RewardTotal=grpc.GetRewardByCustomerId(id_string)
		account.Cart=grpc.GetCartByCustomerId(id_string)
		return 200,account,""
}