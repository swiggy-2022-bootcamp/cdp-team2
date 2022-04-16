package dao

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/errors"

	log "github.com/sirupsen/logrus"
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
		log.Info("Record not found for customer ", reward.CustomerId)
	} else {
		totalPoints := rewardRecord.Points + reward.Points
		reward.Points = totalPoints
	}

	av, goErr := dynamodbattribute.MarshalMap(reward)
	if goErr != nil {
		log.WithError(goErr).Error("an error occurred while marshalling reward to dynamo attribute")
		return &errors.MarshalToAttributesError
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("reward"),
	}

	_, goErr = dao.client.PutItem(input)
	if goErr != nil {
		log.WithField("Error: ", goErr).Error("got error while calling GetItem API of dynamo db")
		return &errors.DatabaseQueryFailed
	}

	log.Info("successfully added the reward points to the customer: ", reward.CustomerId)
	return nil
}

// GetReward gets the reward associated with the customer
func (dao *dynamoDAO) GetReward(customerId string) (models.Reward, *errors.ServerError) {
	reward := models.Reward{}
	result, err := dao.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("reward"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
	})
	if err != nil {
		log.WithField("Error: ", err).Error("got error while calling GetItem API of dynamo db")
		return reward, &errors.DatabaseQueryFailed
	}

	if result.Item == nil {
		errMsg := "Could not find '" + customerId + "'"
		log.WithField("Error: ", errMsg).Error("record not found")
		return reward, &errors.RecordNotFound
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &reward)
	if err != nil {
		log.WithError(err).Error("an error occurred while unmarshalling dynamo attributes to reward struct")
		return reward, &errors.UnmarshalAttributesError
	}

	log.Info("Successfully got the reward for the customer: ", customerId, " reward points: ", reward.Points)
	return reward, nil
}
