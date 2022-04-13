package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBaseRepository struct {
	DataStore *mongo.Database
}

var db *DataBaseRepository

func ConnectDB() *DataBaseRepository {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connecting to MonogDB")
	if db == nil {
		db = &DataBaseRepository{
			DataStore: client.Database("swiggy"),
		}
	}
	return db
}
