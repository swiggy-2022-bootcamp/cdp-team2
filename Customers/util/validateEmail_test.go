package util_test
import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
 	"github.com/aws/aws-sdk-go/service/dynamodb"
 	model "customers/dao/models"
	"customers/util"
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
	Customer model.Customer
	Expected string
	Status int
}



var tests= []Tests{}

func InitTests(){
	tests=append(tests,Tests{model.Customer{Id:"afgiu234rr",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"1234567",Password:"udaykiran",Confirm:"udaykiran"},"error",200})
	tests=append(tests,Tests{model.Customer{Id:"rgtsyur4ee",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"},"error",500})
	tests=append(tests,Tests{model.Customer{Id:"234erfvefj",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"67",Password:"udayiran",Confirm:"udaykiran"},"error",500})
	tests=append(tests,Tests{model.Customer{Id:"23r4ferdfv",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"",Confirm:"udaykiran"},"error",500})
	tests=append(tests,Tests{model.Customer{Id:"123r4345tf",Lastname:"Uday",Firstname:"",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"},"error",500})
 }

func TestValidateEmail(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
	InitTests()
	for _,val:=range tests{
		status:=util.ValidateEmail(val.Customer.Email,mockSvc)
		if status!=true{
			return 
		}
	}
}