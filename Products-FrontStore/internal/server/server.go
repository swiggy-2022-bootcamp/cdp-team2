package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/products-frontstore-service/config"
	_ "github.com/products-frontstore-service/docs"
	"github.com/products-frontstore-service/internal/core/ports"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Handlers ports.IHandlers
}

func NewServer(handlers ports.IHandlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

// @title Swagger Products FrontStore Microservice
// @version 1.0
// @description Micorservice for handling frontstore products.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /api/rest/products
func (s *Server) Initialize() {
	server := gin.Default()

	productsRoutes := server.Group("/api/rest/products/")
	productsRoutes.GET("/", s.Handlers.Health)
	productsRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(config.Config["PORT"]))
}
