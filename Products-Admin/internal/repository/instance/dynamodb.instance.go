package instance

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
	Initialize a session that the SDK will use to load
	credentials from the shared credentials file ~/.aws/credentials
	and region from the shared configuration file ~/.aws/config.
	and returns a new DynamoDB client.
*/
func GetDynamoDBClient() *dynamodb.DynamoDB {
	dbSession := session.New(&aws.Config{
		Region: aws.String("us-west-2"),
	})

	return dynamodb.New(dbSession)
}
