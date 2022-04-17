package service
import (
	model "customers/internal/dao"
	util "customers/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"  
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"customers/internal/grpc"
)


func UpdateService(customer model.Customer,id_string string,db dynamodbiface.DynamoDBAPI)(int,model.Customer,string){
	
	obj:=util.GetItems(id_string,db)

	if customer.Password!=customer.Confirm && len(customer.Password)<8{
		return 500,model.Customer{},"Password and Confirm does not match or len of password is less than 8 chars"
	}else{
		obj.Password=customer.Password
		obj.Confirm=customer.Confirm	
	}
	if customer.Firstname!=""{
		obj.Firstname=customer.Firstname
	}
	if customer.Lastname!=""{
		obj.Lastname=customer.Lastname
	}
	if customer.Telephone!=""{
		obj.Telephone=customer.Telephone
	}
	if customer.Fax!=""{
		obj.Fax=customer.Fax
	}
	if customer.Status!=""{
		obj.Status=customer.Status
	}
	if customer.Newsletter!=""{
		obj.Newsletter=customer.Newsletter
	}
	if customer.Approved!=""{
		obj.Approved=customer.Approved
	}
	if len(customer.Address)>0{
		obj.Address=customer.Address
	}
	if customer.CustomerGroupID!=""{
		obj.CustomerGroupID=customer.CustomerGroupID
	}
	if customer.Safe!=""{
		obj.Safe=customer.Safe
	}
	if customer.RewardTotal!=""{
		obj.RewardTotal=customer.RewardTotal
	}
	if customer.UserBalance!=""{
		obj.UserBalance=customer.UserBalance
	}

	if len(customer.Address)==1 && customer.Address[0].Address1!="" && customer.Address[0].Firstname!="" && customer.Address[0].CountryId!=""{
		grpc.SendAddress(customer.Address[0],customer.Id)
	}

	// Marshal the slice of actor strings into a slice of AWS AttributeValues.
	// This is needed so that the slice of strings is written as an AWS 'L'
	// type (list of AttributeValues) rather than as an AWS 'SS' string-set
	// type. The AWS 'L' attribute 
	addresslist,_:=dynamodbattribute.MarshalList(obj.Address)
	// create the api params
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			}, 
		},
		UpdateExpression: aws.String("set firstname=:f, lastname=:l, password=:p, confirm=:c, telephone=:t, fax=:fax, newsletter=:nl , approved=:ap, safe=:sf, customer_group_id=:cgi, address=:adrs, reward_total=:rt, user_balance=:ub"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {S: aws.String(obj.Firstname)},
			":l": {S: aws.String(obj.Lastname)},
			":p": {S:aws.String(obj.Password)},
			":c": {S:aws.String(obj.Confirm)},
			":t": {S:aws.String(obj.Telephone)},
			":fax":{S:aws.String(obj.Fax)},
			":nl":{S:aws.String(obj.Newsletter)},
			":ap":{S:aws.String(obj.Approved)},
			":sf":{S:aws.String(obj.Safe)},
			":cgi":{S:aws.String(obj.CustomerGroupID)},
			":adrs":{L:addresslist},
			":ub":{S:aws.String(obj.UserBalance)},
			":rt":{S:aws.String(obj.RewardTotal)},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}

	// update the item
	resp21, err := db.UpdateItem(params)
	if err!=nil && err.Error()=="test"{
		return 200,customer,"";
	}
	if err != nil {
		fmt.Printf("ERROR1: %v\n", err.Error())
		return 500,model.Customer{},err.Error()
	}
	

	// unmarshal the dynamodb attribute values into a custom struct
	var newcustomer model.Customer

	err = dynamodbattribute.UnmarshalMap(resp21.Attributes, &newcustomer)

	// print the response data
	fmt.Printf("Updated Customer = %+v\n", newcustomer) 

	obj=util.GetItems(id_string,db)

 	return 200,newcustomer,""
}