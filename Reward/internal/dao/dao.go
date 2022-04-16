package dao

import (
	"fmt"
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/errors"
)

type DynamoDAO interface {
	AddReward(reward models.Reward) *errors.ServerError
	GetReward(customerId string) (models.Reward, *errors.ServerError)
}

type dynamoDAO struct {
	client *dynamodb.DynamoDB
	config *config.WebServerConfig
}

var dynamoDao DynamoDAO
var dynamoDaoOnce sync.Once

func InitDynamoDAO(client *dynamodb.DynamoDB, config *config.WebServerConfig) DynamoDAO {
	dynamoDaoOnce.Do(func() {
		dynamoDao = &dynamoDAO{
			client: client,
			config: config,
		}
	})

	return dynamoDao
}

func GetDynamoDAO() DynamoDAO {
	if dynamoDao == nil {
		panic("Dynamo DAO not initialised")
	}

	return dynamoDao
}

// AddProduct add product details to the cart of the user
func (dao *dynamoDAO) AddReward(reward models.Reward) *errors.ServerError {
	rewardRecord, err := dao.GetReward(reward.CustomerId)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	totalPoints := rewardRecord.Points + reward.Points
	reward.Points = totalPoints

	av, goErr := dynamodbattribute.MarshalMap(reward)
	if goErr != nil {
		log.Fatalf("Got error marshalling new movie item: %s", goErr)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("reward"),
	}

	_, goErr = dao.client.PutItem(input)
	if goErr != nil {
		log.Fatalf("Got error calling PutItem: %s", goErr)
	}

	fmt.Println("Successfully added ")
	return nil
}

// GetReward gets the reward associated with the customer
func (dao *dynamoDAO) GetReward(customerId string) (models.Reward, *errors.ServerError) {
	reward := models.Reward{}
	result, err := dao.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("reward"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				N: aws.String(customerId),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find '" + customerId + "'"
		log.Fatal(msg)
		return reward, nil
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &reward)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Customer Id:  ", reward.CustomerId)
	fmt.Println("Description: ", reward.Description)
	fmt.Println("Points:  ", reward.Points)

	return reward, nil
}
