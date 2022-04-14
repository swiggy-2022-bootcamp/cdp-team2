package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/products-admin-service/internal/core/ports"
	"github.com/products-admin-service/internal/core/services"
	"github.com/products-admin-service/internal/handlers"
	"github.com/products-admin-service/internal/repository"
	"github.com/products-admin-service/internal/server"
)

var (
	err                error
	productsServer     ports.IServer
	productsHandlers   ports.IProdcutsHandlers
	productsServices   ports.IProductsServices
	productsRepository ports.IProductsRepository
)

func init() {
	// Load Dotenv file
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error Loading in .env file: ", err.Error())
	}

	// InitialiZe Repository
	// TODO: Need to load .env file via config.
	productsRepository, err = repository.NewProductsRepository(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err)
	}

	productsRepository, _ = repository.NewProductsRepository("dummy")
	productsServices = services.NewProductsServices(productsRepository)
	productsHandlers = handlers.NewHandlers(productsServices)
	productsServer = server.NewServer(productsHandlers)
}

func main() {
	productsServer.Initialize()
}
