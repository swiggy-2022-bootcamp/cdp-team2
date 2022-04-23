package util_test
import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
 	"github.com/aws/aws-sdk-go/service/dynamodb"
 	model "customer-account/dao/models"
	"github.com/stretchr/testify/assert"
	"customer-account/util"
	"fmt"
)

type mockDynamoDBClient struct {
    dynamodbiface.DynamoDBAPI
}

func  (m *mockDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error){
	return &dynamodb.PutItemOutput{},nil
}
func    (m *mockDynamoDBClient)  Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error){
	return &dynamodb.QueryOutput{},nil
}

type Tests struct{
	Account model.Account
	
}



var tests= []Tests{}

func InitTests(){
	tests=append(tests,Tests{model.Account{Id:"afgiu234rr",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"1234567",Password:"udaykiran",Confirm:"udaykiran"}})
	tests=append(tests,Tests{model.Account{Id:"rgtsyur4ee",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"}})
	tests=append(tests,Tests{model.Account{Id:"234erfvefj",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"67",Password:"udayiran",Confirm:"udaykiran"}})
	tests=append(tests,Tests{model.Account{Id:"23r4ferdfv",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"",Confirm:"udaykiran"}})
	tests=append(tests,Tests{model.Account{Id:"123r4345tf",Lastname:"Uday",Firstname:"",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"}})
 }

func TestValidateEmail(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
	InitTests()
	fmt.Println("Tested ValidateEmail")

	for _,val:=range tests{
		status:=util.ValidateEmail(val.Account.Email,mockSvc)
		assert.Equal(t,true,status)
	}
}