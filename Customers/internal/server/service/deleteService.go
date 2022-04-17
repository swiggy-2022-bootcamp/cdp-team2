package service
import (
 	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"fmt" 
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

)

func DeleteService(id_string string,db dynamodbiface.DynamoDBAPI)(int,string){
	fmt.Println("delete:----",id_string)
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(id_string),
			},
		 
		},
	}

	// delete the item
	_, err := db.DeleteItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return 500,err.Error()
	}

	// print the response data
	 
	return 200,"Deleted Successfully!1"
}