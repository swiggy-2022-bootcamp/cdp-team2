package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)


var Server = map[string]string{
	"PORT": "",
	"CART_PORT":"",
	"REWARD_PORT":"",
	"AUTH_PORT":"",
	"GRPC_PORT":"",
}

var AWS = map[string]string{
	"region":   "",
	"endpoint": "",
	"TABLE_NAME":"",
}



func FromEnv()  {
	cfgFilename := "../../etc/.env"
	if err := godotenv.Load(cfgFilename); err != nil {
		fmt.Println("No config files found to load to env.",err.Error())
	}

	Server["PORT"]="8091"
	//os.Getenv("PORT")
	fmt.Println("__",os.Getenv("PORT"))

	Server["CART_PORT"]="9000"
	//os.Getenv("CART_PORT")
	Server["REWARD_PORT"]="9000"
	//os.Getenv("REWARD_PORT")
	Server["GRPC_PORT"]="9000"
	//os.Getenv("GRPC_PORT")
	Server["AUTH_PORT"]="9000"
	//os.Getenv("AUTH_PORT")

	AWS["TABLE_NAME"]="team-2-Customers"
	//os.Getenv("TABLE_NAME")
	AWS["region"]="us-west-2"
	//os.Getenv("REGION")
	AWS["endpoint"]="https://dynamodb.us-west-2.amazonaws.com"
	//os.Getenv("ENDPOINT")
	
}

// package main

// import (
// 	"fmt"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// )
// var Database  *dynamodb.DynamoDB
// func InitDB(){
// 		// create an aws session
// 		sess := session.Must(session.NewSession(&aws.Config{
// 			Region:   aws.String("us-west-2"),

// 		//	Endpoint: aws.String("http://127.0.0.1:8000"),
// 		}))
	
// 		// create a dynamodb instance
// 		db := dynamodb.New(sess,&aws.Config{Endpoint:aws.String("https://dynamodb.us-west-2.amazonaws.com"),})
// 	//	CreateTable(db)
// 		Database=db
// }

// func CreateTable(db *dynamodb.DynamoDB){
	
// 	// create the api params
// 	params := &dynamodb.CreateTableInput{
// 		TableName: aws.String("team-2"),
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{AttributeName: aws.String("customer_id"), KeyType: aws.String("HASH")},
// 		//	{AttributeName: aws.String("email"), KeyType: aws.String("RANGE")},
// 		},
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{AttributeName: aws.String("customer_id"), AttributeType: aws.String("S")},
// 	//		{AttributeName: aws.String("email"), AttributeType: aws.String("S")},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(1),
// 			WriteCapacityUnits: aws.Int64(1),
// 		},
// 	}

// 	// create the table
// 	resp, err := db.CreateTable(params)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// print the response data
// 	fmt.Println(resp)
// }
// func main(){
// 	InitDB()
// 	CreateTable(Database)
// }
// func GetDB()*dynamodb.DynamoDB{
// 	fmt.Println(Database)
// 	return Database
// }