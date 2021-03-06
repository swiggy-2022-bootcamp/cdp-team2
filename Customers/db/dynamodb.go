package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "customers/config"
)

var db *dynamodb.DynamoDB

//Returns instance of DynamoDB
//Creates instance if not already made
//go:generate mockgen --destination=../mocks/mock_dynamodbiface/dynamodbiface.go github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface DynamoDBAPI
func GetInstance() *dynamodb.DynamoDB {
	if db == nil {
		// db = dynamodb.New(session.New(&aws.Config{
		// 	Region: aws.String(config.AWS["region"]),
		// 	// Credentials: credentials.NewEnvCredentials(),
		// 	Endpoint: aws.String(config.AWS["endpoint"]),
		// }))
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
