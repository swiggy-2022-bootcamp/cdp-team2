package service
import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
 	"github.com/aws/aws-sdk-go/service/dynamodb"
 	model "customer-account/internal/dao"
	
	"errors"
)

type mockDynamoDBClient struct {
    dynamodbiface.DynamoDBAPI
}

func  (m *mockDynamoDBClient) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error){
	return nil,errors.New("test")
}
// func    (m *mockDynamoDBClient)  Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error){
// 	return nil,errors.New("Test")
// }

type Tests struct{
	Customer model.Account
	Expected string
	Status int
}

var tests= []Tests{}

func InitTests(){
	tests=append(tests,Tests{model.Account{Id:"afgiu234rr",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"1234567",Password:"udaykiran",Confirm:"udaykiran"},"error",200})
	tests=append(tests,Tests{model.Account{Id:"rgtsyur4ee",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"},"error",500})
	tests=append(tests,Tests{model.Account{Id:"234erfvefj",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"67",Password:"udayiran",Confirm:"udaykiran"},"error",500})
	tests=append(tests,Tests{model.Account{Id:"23r4ferdfv",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"",Confirm:"udaykiran"},"error",500})
	tests=append(tests,Tests{model.Account{Id:"123r4345tf",Lastname:"Uday",Firstname:"",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"},"error",500})
}

func TestCreateService(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
	InitTests()
	for _,val:=range tests{
		status,_,_:=CreateService(val.Customer,mockSvc)
 		if(val.Status!=status){
			t.Errorf("Error returned: %v, want: %v", status,val.Status)
		}
	}
	got:=1
	want:=1
	if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

}