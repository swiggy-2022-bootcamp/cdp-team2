package service
import (
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
 	"github.com/aws/aws-sdk-go/service/dynamodb"
	"errors"
 )
 
 

func  (m *mockDynamoDBClient) UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error){

	return nil,errors.New("test")
}
  
func   (m *mockDynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error){
	return nil,errors.New("test")
 }

 
func TestUpdateService(t *testing.T){
	mockSvc := &mockDynamoDBClient{}

	for _,val := range tests{
		UpdateService(val.Customer,val.Customer.Id,mockSvc) 
		 
	}
	
}