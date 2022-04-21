package service

import (
 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	util "customer-account/util"
 	"customer-account/dao"
	"customer-account/dao/models"
	"errors"
)


func ValidateAccount(account models.Account,db dynamodbiface.DynamoDBAPI)(error){
 
 
	if len(account.Firstname)<=1{
		return errors.New("Please enter a valid Firstname")
	}
	if len(account.Lastname)<=1{
		return errors.New("Please enter a valid Lastname")
	}
	if util.ValidateEmail(account.Email,db)==false{
		return errors.New("Email already exists")
	} else if len(account.Email)<5{
		return errors.New("Please Enter a valid email")
	}
	if account.Password!=account.Confirm{
		return errors.New("Password and Confirm does not match")
	} else if len(account.Password)<8{
		return errors.New("Password should be minimum 8 characters")
	}
	if len(account.Telephone)<=4{
		return errors.New("Please enter a valid phone number")
	}

	return nil
	
}

type CustomerService struct{
	Dao dao.IDao
} 

func NewCustomerService() IService{
	return &CustomerService{
		dao.GetCustomerDao(),
	}
}

func (cs *CustomerService)CreateService(account models.Account)(models.Account,error){
	return dao.GetCustomerDao().Create(account)
}



func (cs *CustomerService)ReadService(id_string string)( models.Account2,error){
	return dao.GetCustomerDao().Get(id_string)
}


func  (cs *CustomerService)UpdateService(id_string string,customer models.Account)(models.Account,error){
	return dao.GetCustomerDao().Update(id_string,customer)
}