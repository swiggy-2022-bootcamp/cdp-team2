package main

import (
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/services"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/handlers"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/repository/adaptor"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/repository/instance"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/server"
	"google.golang.org/grpc"
)

var (
	err                error
	productsServer     ports.IServer
	productsHandlers   ports.IProductsHandlers
	productsServices   ports.IProductsServices
	productsRepository ports.IProductsRepository
	dynamodbClient     *dynamodb.DynamoDB
	gRPCServer         *grpc.Server
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
