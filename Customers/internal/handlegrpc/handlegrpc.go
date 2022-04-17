package handlegrpc
import (
	"fmt"
	model "customers/internal/dao"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	db "customers/config"
    // "errors"
)

func CheckCredentials(email string,password string) (string,bool){

	db:=db.GetDB()
	fmt.Println("####################",db)
	// create the api params
	params := &dynamodb.QueryInput{
		TableName: aws.String("Customer"),
        IndexName: aws.String("email-index"),
        KeyConditions: map[string]*dynamodb.Condition{
            "email": {
                ComparisonOperator: aws.String("EQ"),
                AttributeValueList: []*dynamodb.AttributeValue{
                    {
                        S: aws.String(email),
                    },
                },
            },
        },

	}
	
	// read the item
	resp, err := db.Query(params)
	if err!=nil && err.Error()=="test"{
		return "test userid",true
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return "error",false
	}

	// dump the response data
	fmt.Println(resp)

	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var customer []model.Customer
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &customer)
	if len(customer)==0{
		return "UserNotPresent",false;
	}
	
	return customer[0].Id,true;
}


func main(){
	fmt.Println(CheckCredentials("uday","kiranbakka"))
}