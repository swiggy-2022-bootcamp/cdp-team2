package dao


import (
	"errors"
	"fmt"
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"customers/util"
 	"customers/dao/models"
	"customers/db"
	"customers/internal/grpc"
	// "customers/config"
	"strconv"
	"time"
	"crypto/sha1"
	"encoding/hex"
	log "github.com/sirupsen/logrus"

)

var (
	tableName = "team-2-Customers"
)

type CustomerDao struct{
	Db dynamodbiface.DynamoDBAPI
}

func GetCustomerDao() IDao{
	return &CustomerDao{
		db.GetInstance(),
	}
}

func customersTableName() *string{
	return aws.String(tableName)
}
 

func (cd *CustomerDao) Create(customer models.Customer) (models.Customer, error) {
	account:=customer
 	err:=util.ValidateAccount(account,cd.Db)
  	if err!=nil{
		log.WithField("Error: ", err).Error("Error while creating the user")
 		return models.Customer{},err
	}
	if customer.Id!="test"{
		customer.Id=util.Generate()
	}
	if customer.DateAdded!="test"{
		customer.DateAdded=time.Now().String()
	}

	if customer.Id!="test"{
		h := sha1.New()
		h.Write([]byte(customer.Password))
		customer.Password = hex.EncodeToString(h.Sum(nil))
	
		h = sha1.New()
		h.Write([]byte(customer.Confirm))
		customer.Confirm = hex.EncodeToString(h.Sum(nil))
	}
	
	
	info, err := dynamodbattribute.MarshalMap(customer)
 
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the movie, %v", err))
		return models.Customer{},err
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String(tableName),
	}

	_, err = cd.Db.PutItem(input)
 	if  len(customer.Address)==1 && customer.Address[0].AddressLine1!="" && customer.Address[0].Firstname!="" && customer.Address[0].CountryCode!=""{
		grpc.SendAddress(customer.Address[0],customer.Id)

 	}
	 
	if err!= nil {
		log.WithField("Error: ", err).Error("Error while creating ther customer #2")
		return models.Customer{},err
	}
	customer.Address=grpc.GetAddress(customer.Id)
 	return customer,nil
	
}

func Validate(obj models.Customer,customer models.Customer)(models.Customer,error){
	
	if customer.Password!=customer.Confirm && len(customer.Password)<8{
		return models.Customer{},errors.New("Password and Confirm does not match or len of password is less than 8 chars")
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
		if customer.Status!="1" && customer.Status!="0"{
			return obj,errors.New("Status should be either \"0\" or \"1\"")
		}
		obj.Status=customer.Status
	}
	if customer.Newsletter!=""{
		obj.Newsletter=customer.Newsletter
	}
	if customer.Approved!=""{
		if customer.Approved!="1" && customer.Approved!="0"{
			return obj,errors.New("Approved should be either \"0\" or \"1\"")
		}	
		obj.Approved=customer.Approved
	}
	if len(customer.Address)>0{
		obj.Address=customer.Address
 		grpc.SendAddress(customer.Address[0],obj.Id)
	}
	if customer.CustomerGroupID!=0{
		obj.CustomerGroupID=customer.CustomerGroupID
	}
	if customer.Safe!=""{
		if customer.Safe!="1" && customer.Safe!="0"{
			return obj,errors.New("Customer Safe should be either \"0\" or \"1\"")
		}
		obj.Safe=customer.Safe
	}
 
	
	return obj,nil
}

func (cd *CustomerDao)Update(id_string string,customer models.Customer)(models.Customer,error){

	obj:=util.GetItems(id_string,cd.Db)
	if obj.Id==""{
		return obj,errors.New("Invalid User id")
	}
 
	var err1 error
	obj,err1=Validate(obj,customer)
 	if err1!=nil{
		log.WithField("Error: ", err1).Error("Error while Updating the Customer")
		return obj,err1
	}

	if customer.Password!=customer.Confirm && len(customer.Password)<8{
		return models.Customer{},errors.New("Password and Confirm does not match or len of password is less than 8 chars")
	} else if len(customer.Password)>=8{

		if(customer.Password!=customer.Confirm){
			return models.Customer{},errors.New("Password and Confirm does not match or len of password is less than 8 chars")
		}
		h := sha1.New()
		h.Write([]byte(customer.Password))
		obj.Password = hex.EncodeToString(h.Sum(nil))
		h = sha1.New()
		h.Write([]byte(customer.Confirm))
		obj.Confirm = hex.EncodeToString(h.Sum(nil))
		customer.Password=obj.Password
		customer.Confirm=obj.Confirm
	}else{
		obj.Password=customer.Password
		obj.Confirm=customer.Confirm	
	}
 

	 


	if len(customer.Address)==1 && customer.Address[0].AddressLine1!="" && customer.Address[0].Firstname!="" && customer.Address[0].CountryCode!=""{
		grpc.SendAddress(customer.Address[0],customer.Id)
	}

	// Marshal the slice of actor strings into a slice of AWS AttributeValues.
	// This is needed so that the slice of strings is written as an AWS 'L'
	// type (list of AttributeValues) rather than as an AWS 'SS' string-set
	// type. The AWS 'L' attribute 
	addresslist,_:=dynamodbattribute.MarshalList(obj.Address)
	// create the api params
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			}, 
		},
		UpdateExpression: aws.String("set firstname=:f, lastname=:l, password=:p, confirm=:c, telephone=:t, fax=:fax, newsletter=:nl , approved=:ap, safe=:sf, customer_group_id=:cgi, address=:adrs, statuss=:stas "),
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
			":cgi":{N:aws.String(strconv.Itoa(obj.CustomerGroupID))},
			":adrs":{L:addresslist},
			":stas":{S:aws.String(obj.Status)},
 		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}

	// update the item
	resp, err := cd.Db.UpdateItem(params)
 
	if err != nil {
		log.WithField("Error: ", err).Error("error while updating the customer")

		return models.Customer{},err
	}
	

	// unmarshal the dynamodb attribute values into a custom struct
	var newcustomer models.Customer

	err = dynamodbattribute.UnmarshalMap(resp.Attributes, &newcustomer)

	// print the response data

	obj=util.GetItems(id_string,cd.Db)

 	return newcustomer,nil
}


func (cd *CustomerDao)Get(id_string string)(models.Customer,error){
	params := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
	 
			"customer_id": {
				S: aws.String(id_string),
			},
 
		},
	}
	var customer models.Customer
 	// read the item
	resp, err := cd.Db.GetItem(params)
 	if err != nil {
		log.WithField("Error: ", err).Error("err while getting Email")
 		return models.Customer{},nil
	}
 	// dump the response data
 	// unmarshal the dynamodb attribute values into a custom struct
	err = dynamodbattribute.UnmarshalMap(resp.Item, &customer)
 	if err != nil {
		log.WithField("Error: ", err).Error("Error while getting user")

 		return models.Customer{},err
	}
	if customer.Id==""{
		return customer,errors.New("Customer Not Found")
	}
	// print the response data
 	customer.Address=grpc.GetAddress("0")
	
	 return customer,nil
}

func (cd *CustomerDao)GetByEmail(email string)(models.Customer,error){
	params := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
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
	resp, err := cd.Db.Query(params)
 
	if err != nil {
		log.WithField("Error: ", err).Error("Error while getting the email")

 		return models.Customer{},err
	}

	// dump the response data
 
	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var account []models.Customer
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &account)
	
	
	
	if len(account)==0 {
		return models.Customer{},errors.New("Customer with this email does not exist!!");
	}

	account[0].Address=grpc.GetAddress(account[0].Id)
	
	return account[0],nil
}

func (cd *CustomerDao)Delete(id_string string)(bool,error){
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			},
		 
		},
	}

	// delete the item
	_, err := cd.Db.DeleteItem(params)
	if err != nil {
		log.WithField("Error: ", err).Error("Error while deleting the customer")
 		return false,err
	}

	 
	return true,nil

}
