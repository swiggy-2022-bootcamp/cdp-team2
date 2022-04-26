package repository

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Access Key ID:
// AKIAQ6LFILQYZPAAQPGT
// Secret Access Key:
// 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
// AWS Account ID: 065173740593
// Canonical User ID: 42102cdb0b0f0278b803eb3d703cb42515d4a5b4dbd72f381f14e954ca4511da

var Database *dynamodb.DynamoDB

func InitDB() {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),

		//	Endpoint: aws.String("http://127.0.0.1:8000"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("https://dynamodb.us-east-1.amazonaws.com")})
	fmt.Println(reflect.TypeOf(db))
	CreateTable(db)
	Database = db
}

func CreateTable(db *dynamodb.DynamoDB) {

	// create the api params
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Customers"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("customer_id"), KeyType: aws.String("HASH")},
			//	{AttributeName: aws.String("email"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("customer_id"), AttributeType: aws.String("S")},
			//		{AttributeName: aws.String("email"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	// create the table
	resp, err := db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the response data
	fmt.Println(resp)
}

func GetDB() *dynamodb.DynamoDB {

	return Database
}
