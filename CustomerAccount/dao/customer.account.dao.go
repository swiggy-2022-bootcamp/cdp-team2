package dao


import (
	"errors"
	"fmt"
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"customer-account/util"
 	"customer-account/dao/models"
	"customer-account/db"
	"customer-account/internal/grpc"
	"time"
	log "github.com/sirupsen/logrus"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
	"customer-account/config"
)

var (
	tableName = config.AWS["TABLE_NAME"]
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
	return aws.String("team-2-Customers")
}
 
func (cd *CustomerDao) Create(account models.Account) (models.Account, error) {
	if(account.Id!="test"){
		account.Id = util.Generate()
	}
	
	err:=util.ValidateAccount(account,cd.Db);

	if err!=nil{
		return models.Account{},err;
	}
	if account.DateAdded!="test"{
		account.DateAdded=time.Now().String()
	}
	info, err := dynamodbattribute.MarshalMap(account)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal the movie, %v", err))
		return models.Account{},err
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("team-2-Customers"),
	}

	_, err = cd.Db.PutItem(input)

	if err != nil {
 		return models.Account{},err
	}
 
	return account,nil
	
}


func (cd *CustomerDao)Update(id_string string,customer models.Account)(models.Account,error){

	obj:=util.GetItems(id_string,cd.Db)

	if customer.Password!=customer.Confirm && len(customer.Password)<8{
		return models.Account{},errors.New("Password and Confirm does not match or len of password is less than 8 chars")
	}else{
		obj.Password=customer.Password
		obj.Confirm=customer.Confirm	
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
	if obj.CustomerGroupID!=customer.CustomerGroupID{
		obj.CustomerGroupID=customer.CustomerGroupID
	}
	if obj.UserBalance!=customer.UserBalance{
		obj.UserBalance=customer.UserBalance
	}
	if len(customer.Email)>0{
 
			return models.Account{},errors.New("Please don't change email!")
	
	}
	 
	
	h := sha1.New()
	h.Write([]byte(obj.Password))
	password_hash := hex.EncodeToString(h.Sum(nil))


	h = sha1.New()
	h.Write([]byte(obj.Confirm))
	confirm_hash := hex.EncodeToString(h.Sum(nil))

	
	// create the api params
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("team-2-Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			}, 
		},
		UpdateExpression: aws.String("set firstname=:f, lastname=:l, password=:p, confirm=:c, telephone=:t, email=:e, user_balance=:ub, customer_group_id=:cgi"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {S: aws.String(obj.Firstname)},
			":l": {S: aws.String(obj.Lastname)},
			":p": {S:aws.String(password_hash)},
			":c": {S:aws.String(confirm_hash)},
			":t": {S:aws.String(obj.Telephone)},
			":e": {S:aws.String(obj.Email)},
			":cgi":{N:aws.String(strconv.Itoa(obj.CustomerGroupID))},
			":ub" :{N:aws.String(strconv.FormatFloat(obj.UserBalance, 'E', -1, 32))},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}

	// update the item
	resp, err := cd.Db.UpdateItem(params)
 
	if err!= nil {
		fmt.Printf("ERROR1: %v\n", err.Error())
		return models.Account{},err
	}

	// unmarshal the dynamodb attribute values into a custom struct
	var newaccount models.Account

	err = dynamodbattribute.UnmarshalMap(resp.Attributes, &newaccount)

	// print the response data
	log.WithField("Error:",err).Error("Error while Creating User")
	obj=util.GetItems(id_string,cd.Db)

	return obj,nil;
}


func (cd *CustomerDao)Get(id_string string)(models.Account2,error){
	fmt.Println("table name",id_string,"team-2-Customers")
	params := &dynamodb.GetItemInput{
		TableName: aws.String("team-2-Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			},
		},
	}

	resp, err := cd.Db.GetItem(params)
	log.WithField("Error:",err).Error("encountered whiel")
	fmt.Println("----------------",err,resp)
	if err != nil {
		// log.Error("ERROR: %v\n", err.Error())
		return models.Account2{},err
	}
 
  
	// unmarshal the dynamodb attribute values into a custom struct
	var account models.Account2
	err = dynamodbattribute.UnmarshalMap(resp.Item, &account)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return models.Account2{},err
	}
	if(id_string=="test"){
		return account,nil;
	}
 
	if(account.Id==""){
		return account,errors.New("User not found");
	}

	// print the response data
	fmt.Printf("Unmarshaled Movie-- = %+v\n", account)
	
	account.Reward=grpc.GetRewardByCustomerId(id_string)
	account.Cart=grpc.GetCartByCustomerId(id_string)
	return account,nil

}