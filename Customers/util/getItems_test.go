package util
import (
 	"testing"
	 "github.com/aws/aws-sdk-go/service/dynamodb"
 	"errors"
 )

 
 
 func   (m *mockDynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error){
	return nil,errors.New("test")
 }

func TestGetItems(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
 	for _,val:=range tests{
		status:=GetItems(val.Customer.Id,mockSvc)
		if status.Id!="test"{
			t.Errorf("Did not returned the Customer %v",status)
		}
	}
}