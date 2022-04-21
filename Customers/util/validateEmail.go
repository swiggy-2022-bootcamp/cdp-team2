package util


import (
 	model "customers/dao/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func ValidateEmail(email string,db dynamodbiface.DynamoDBAPI) bool{
	
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
 
 	if err != nil {
 		return false;
	}

	// dump the response data

	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var customer []model.Customer
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &customer)
	if len(customer)>0{
		return false;
	}
	return true;
}


