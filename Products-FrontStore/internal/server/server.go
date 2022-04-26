package server

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/config"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/ports"
)

type Server struct {
	Handlers ports.IProductsHandlers
}

func NewServer(handlers ports.IProductsHandlers) *Server {
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

// @host 35.84.28.237:30203
// @BasePath /api/rest/products
func (s *Server) Initialize() {
	server := gin.Default()
	productsRoutes := server.Group("/api/rest/products/")
	productsRoutes.GET("/health", s.Handlers.Health)
	productsRoutes.GET("/", s.Handlers.GetProductList)
	productsRoutes.GET("/:id", s.Handlers.GetProductById)
	productsRoutes.GET("/category/:id", s.Handlers.GetProductListByCategoryId)

	productsRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(config.Config["PORT"]))
}
