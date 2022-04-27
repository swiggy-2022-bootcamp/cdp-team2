package util_test
import (
 	"testing"
	 "github.com/aws/aws-sdk-go/service/dynamodb"
 	"github.com/stretchr/testify/assert"
	 "customer-account/dao/models"
	 "customer-account/util"
"fmt"
 )

 
 
 func   (m *mockDynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error){
	return &dynamodb.GetItemOutput{},nil
 }

func TestGetItems(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
	fmt.Println("Tested GetItems")

 	for _,val:=range tests{
		status:=util.GetItems(val.Account.Id,mockSvc)
		assert.Equal(t,status,models.Account{})
	}
}