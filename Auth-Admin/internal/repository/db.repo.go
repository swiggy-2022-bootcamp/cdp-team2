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
	Context   *context.Context
	Client    *mongo.Client
}

var db *DataBaseRepository

func ConnectDB() *DataBaseRepository {
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://Khush:1hsgvKPptbsLHPdv@cluster0.87kgy.mongodb.net/swiggy?retryWrites=true"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connecting to MonogDB")
	if db == nil {
		db = &DataBaseRepository{
			DataStore: client.Database("swiggy"),
			Context:   &ctx,
			Client:    client,
		}
	}
	return db
}
