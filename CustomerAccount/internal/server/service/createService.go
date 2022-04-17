package service

import (
 	model "customer-account/internal/dao"
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	util "customer-account/util"
	"time"

)
func ValidateAccount(account model.Account,db dynamodbiface.DynamoDBAPI)(int,string){

	if len(account.UserBalance)==0{
		account.UserBalance="$0.00"
	}
 
	if len(account.Firstname)<=1{
		return 500,"Please enter a valid Firstname"
	}
	if len(account.Lastname)<=1{
		return 500,"Please enter a valid Lastname"
	}
	if util.ValidateEmail(account.Email,db)==false{
		return 500,"Email already exists"
	} else if len(account.Email)<5{
		return 500,"Please Enter a valid email"
	}
	if account.Password!=account.Confirm{
		return 500,"Password and Confirm does not match"
	} else if len(account.Password)<8{
		return 500,"Password should be minimum 8 characters"
	}
	if len(account.Telephone)<=4{
		return 500,"Please enter a valid phone number"
	}

	return 200,"Everything is fine";
	
}

func CreateService(account model.Account,db dynamodbiface.DynamoDBAPI)(int,model.Account,string){
	
	status,message:=ValidateAccount(account,db);
	if status!=200{
		return status,model.Account{},message;
	}
	account.DateAdded=time.Now().String()
	info, err := dynamodbattribute.MarshalMap(account)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the movie, %v", err))
		return 500,model.Account{},err.Error()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Customer"),
	}

	_, err2 := db.PutItem(input)
	if err2!=nil && err2.Error()=="test"{
		return 200,model.Account{},err2.Error()
	}

	if err2 != nil {
		fmt.Println(err2)
		return 500,model.Account{},err2.Error()
	}
	fmt.Println("&&&&&&&&&&&&&&&",account)
	return 200,account,"";
}


