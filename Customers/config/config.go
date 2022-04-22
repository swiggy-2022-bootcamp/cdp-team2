package config
// import (
// 	"fmt"
// )


var Server = map[string]string{
	"PORT": "8050",
}

var AWS = map[string]string{
	"region":   "us-east-1",
	"endpoint": "https://dynamodb.us-east-1.amazonaws.com",
}
 

// var Database  *dynamodb.DynamoDB
// func InitDB(){
// 		// create an aws session
// 		sess := session.Must(session.NewSession(&aws.Config{
// 			Region:   aws.String("us-east-1"),

// 		//	Endpoint: aws.String("http://127.0.0.1:8000"),
// 		}))
	
// 		// create a dynamodb instance
// 		db := dynamodb.New(sess,&aws.Config{Endpoint:aws.String("https://dynamodb.us-east-1.amazonaws.com"),})
// 		fmt.Println(reflect.TypeOf(db))
// 	//	CreateTable(db)
// 		Database=db
// }

// func CreateTable(db *dynamodb.DynamoDB){
	
// 	// create the api params
// 	params := &dynamodb.CreateTableInput{
// 		TableName: aws.String("Customer"),
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

// func GetDB()*dynamodb.DynamoDB{
// 	fmt.Println(Database)
// 	return Database
// }