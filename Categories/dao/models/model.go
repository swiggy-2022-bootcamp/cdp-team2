package models

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Interface of Models
type IModel interface {
	GetTableName(int) *string
	GetKeyFilter() map[string]*dynamodb.AttributeValue
}
