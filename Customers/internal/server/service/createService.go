package service

import (
	model "customers/internal/dao"
	util "customers/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
 	"time"
	 "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	 "customers/internal/grpc"
)

func ValidateAccount(account model.Customer,db dynamodbiface.DynamoDBAPI)(int,string){
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
	return 200,"Everythng is fine"
}

func CreateService(customer model.Customer,db dynamodbiface.DynamoDBAPI)(int,model.Customer,string){
 
	account:=customer

	status,message:=ValidateAccount(account,db)
	if status!=200{
		return status,model.Customer{},message
	}

	if len(customer.UserBalance)==0{
		customer.UserBalance="$0.00"
	}
	if len(customer.RewardTotal)==0{
		customer.RewardTotal="0"
	}

	customer.DateAdded=time.Now().String()
	
	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the movie, %v", err))
		return 500,model.Customer{},err.Error()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Customer"),
	}

	_, err2 := db.PutItem(input)
	fmt.Println(customer)
	if  len(customer.Address)==1 && customer.Address[0].Address1!="" && customer.Address[0].Firstname!="" && customer.Address[0].CountryId!=""{
		grpc.SendAddress(customer.Address[0],customer.Id)
	}
	if err2 != nil {
		fmt.Println(err2)
		return 500,model.Customer{},err2.Error()
	}
	return 200,customer,"Created";
}