package main

import (
	"log"
	// "os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
	"github.com/products-admin-service/internal/core/ports"
	"github.com/products-admin-service/internal/core/services"
	"github.com/products-admin-service/internal/handlers"
	"github.com/products-admin-service/internal/repository/adaptor"
	"github.com/products-admin-service/internal/repository/instance"
	"github.com/products-admin-service/internal/server"
)

var (
	err                error
	productsServer     ports.IServer
	productsHandlers   ports.IProdcutsHandlers
	productsServices   ports.IProductsServices
	productsRepository ports.IProductsRepository
	dynamodbClient     *dynamodb.DynamoDB
)

func init() {
	// Load Dotenv file
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	// Initialize dynamodb client
	dynamodbClient = instance.GetDynamoDBClient()

	// InitialiZe Repository
	// TODO: Need to load .env file via config.
	// productsRepository, err = repository.NewProductsRepository(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}

	// productsRepository, _ = repository.NewProductsRepository("dummy")
	productsRepository = adaptor.NewProductsRepository(dynamodbClient, "team-2-products")
	productsServices = services.NewProductsServices(productsRepository)
	productsHandlers = handlers.NewHandlers(productsServices)
	productsServer = server.NewServer(productsHandlers)
}

func main() {
	productsServer.Initialize()
}
