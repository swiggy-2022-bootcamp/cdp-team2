package service
import (
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
 	"github.com/aws/aws-sdk-go/service/dynamodb"
 	
	// "errors"
)
 
 

func  (m *mockDynamoDBClient)  DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error){
	return nil,nil
}
 
 
func TestDeleteService(t *testing.T){
	mockSvc := &mockDynamoDBClient{}

	status,_:=DeleteService("testId",mockSvc)
	if(status!=200){
			t.Errorf("Error returned: %v, want: %v", status,200)
	}
	
}