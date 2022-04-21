package util_test
import (
 	"testing"
	 "github.com/aws/aws-sdk-go/service/dynamodb"
 	"errors"
	 "customers/dao/models"
	 "github.com/stretchr/testify/assert"
	 "customers/util"
 )

 
 
 func   (m *mockDynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error){
	return nil,errors.New("test")
 }

func TestGetItems(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
 	for _,val:=range tests{
		response:=util.GetItems(val.Customer.Id,mockSvc)
		assert.Equal(t,response,models.Customer{})
	}
}