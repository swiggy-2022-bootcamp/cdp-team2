package util
import (
 	"testing"
  )

 

func TestGenerate(t *testing.T){
	InitTests()
 
		response:=Generate()
		if len(response)!=20{
			t.Errorf("Length generated by this function is not 20 chars")
		}
 
}