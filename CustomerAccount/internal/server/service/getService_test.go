package service
import (
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"errors"
)
 
func    (m *mockDynamoDBClient)  Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error){
	return nil,errors.New("test")
}
 
 
 
func TestReadService(t *testing.T){
	mockSvc := &mockDynamoDBClient{}

	ReadService("testId",mockSvc)
	 
	
}