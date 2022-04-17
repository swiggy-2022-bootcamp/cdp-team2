package service
import (
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
  	
	// "errors"
)
 
 
 
 
func TestReadService(t *testing.T){
	mockSvc := &mockDynamoDBClient{}

	ReadService("testId",mockSvc)
	 
	
}