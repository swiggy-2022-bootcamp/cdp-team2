package service
import (
 	model "customer-account/internal/dao"
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	util "customer-account/util"
)

func UpdateService(id_string string,customer model.Account,db dynamodbiface.DynamoDBAPI)(int,model.Account,string){
	obj:=util.GetItems(id_string,db)

	if customer.Password!=customer.Confirm && len(customer.Password)<8{
		return 500,model.Account{},"Password and Confirm does not match or len of password is less than 8 chars"
	}else{
		obj.Password=customer.Password
		obj.Confirm=customer.Confirm	
	}
	if len(customer.UserBalance)==0{
		obj.UserBalance=customer.UserBalance
	}
	if len(customer.Firstname)>2{
		obj.Firstname=customer.Firstname
	}
	if len(customer.Lastname)>2{
		obj.Lastname=customer.Lastname
	}
	if len(customer.Telephone)>4{
		obj.Telephone=customer.Telephone
	}
	if len(customer.Email)>4{
		if util.ValidateEmail(customer.Email,db)==true{
			return 500,model.Account{},"Email Already exists! please try other email!"
		}
		obj.Email=customer.Email
	}
	 
 
	
	// create the api params
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
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

	// update the item
	resp21, err := db.UpdateItem(params)
	if err!=nil && err.Error()=="test"{
		return 200,model.Account{},err.Error()
	}
	if err!=nil && err != nil {
		fmt.Printf("ERROR1: %v\n", err.Error())
		return 500,model.Account{},err.Error()
	}

	// unmarshal the dynamodb attribute values into a custom struct
	var newaccount model.Account

	err = dynamodbattribute.UnmarshalMap(resp21.Attributes, &newaccount)

	// print the response data
	fmt.Printf("Updated Movie = %+v\n", newaccount) 

	obj=util.GetItems(id_string,db)

	return 200,obj,""
}