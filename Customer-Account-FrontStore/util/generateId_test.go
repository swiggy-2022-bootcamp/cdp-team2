package util_test
import (
 	"testing"
	 "customer-account/util"
	 "fmt"
   )

 

func TestGenerate(t *testing.T){
	InitTests()
	fmt.Println("Tested GenerateId")

		response:=util.Generate()
		if len(response)!=20{
			t.Errorf("Length generated by this function is not 20 chars")
		}
 
}