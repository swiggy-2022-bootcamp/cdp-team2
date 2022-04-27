package util_test

import (
	"fmt"
	"customers/dao/models"
	"customers/util"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)


type Test struct{
	Customer models.Customer
	Err error
}



var test= []Test{}

func InitTest(){
	test=append(test,Test{models.Customer{Id:"afgiu234rr",Lastname:"Uday",Firstname:"i",Email:"mocktest@gmail.com",Telephone:"1234567",Password:"udaykiran",Confirm:"udaykiran"},errors.New("Please enter a valid Firstname")})
	test=append(test,Test{models.Customer{Id:"rgtsyur4ee",Lastname:"u",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayiran",Confirm:"udaykiran"},errors.New("Please enter a valid Lastname")})
	test=append(test,Test{models.Customer{Id:"234erfvefj",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"6888887",Password:"udayiran",Confirm:"udaykian"},errors.New("Password and Confirm does not match")})
	test=append(test,Test{models.Customer{Id:"23r4ferdfv",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"134567",Password:"udayra",Confirm:"udayra"},errors.New("Password should be minimum 8 characters")})
	test=append(test,Test{models.Customer{Id:"123r4345tf",Lastname:"Uday",Firstname:"Kiran",Email:"mocktest@gmail.com",Telephone:"67",Password:"udaykiran",Confirm:"udaykiran"},errors.New("Please enter a valid phone number")})
 }

func TestValidateAccount(t *testing.T){
	mockSvc := &mockDynamoDBClient{}
	InitTest()
	fmt.Println("Tested ValidateAccount")
	for _,val:=range test{
		status:=util.ValidateAccount(val.Customer,mockSvc)
 		assert.Equal(t,status,val.Err)
	}
}