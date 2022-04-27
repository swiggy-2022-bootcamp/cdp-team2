package util

import (
  	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"errors"
	"customer-account/dao/models"	
)
func ValidateAccount(account models.Account,db dynamodbiface.DynamoDBAPI)(error){

 
	if len(account.Firstname)<=1{
		return errors.New("Please enter a valid Firstname")
	}
	if len(account.Lastname)<=1{
		return errors.New("Please enter a valid Lastname")
	}
	if ValidateEmail(account.Email,db)==false{
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
