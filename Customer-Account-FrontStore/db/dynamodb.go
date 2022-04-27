package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "customer-account/config"
)

var db *dynamodb.DynamoDB

//Returns instance of DynamoDB
//Creates instance if not already made
//go:generate mockgen --destination=../mocks/mock_dynamodbiface/dynamodbiface.go github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface DynamoDBAPI
func GetInstance() *dynamodb.DynamoDB {
	if db == nil {
		sess := session.New(&aws.Config{
			Region: aws.String("us-west-2"),
		})
		// sess := session.Must(session.NewSessionWithOptions(session.Options{
		// 	SharedConfigState: session.SharedConfigEnable,
		// }))

		// Create DynamoDB client
		db = dynamodb.New(sess)

	}
	return db
}
