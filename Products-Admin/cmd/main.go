package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
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
	// Initialize dynamodb client
	dynamodbClient = instance.GetDynamoDBClient()

	// InitialiZe Repository
	if err != nil {
		log.Fatal(err)
	}

	productsRepository = adaptor.NewProductsRepository(dynamodbClient, "team-2-products")
	productsServices = services.NewProductsServices(productsRepository)
	productsHandlers = handlers.NewHandlers(productsServices)
	productsServer = server.NewServer(productsHandlers)
}

func main() {
	productsServer.Initialize()
}
