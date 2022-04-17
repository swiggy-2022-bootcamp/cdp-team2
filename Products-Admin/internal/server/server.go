package server

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/config"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
)

type Server struct {
	Handlers ports.IProductsHandlers
}

var _ ports.IServer = (*Server)(nil)

func NewServer(handlers ports.IProductsHandlers) *Server {
	return &Server{
		Handlers: handlers,
	}
}

// @title Swagger Products Admin Microservice
// @version 1.0
// @description Micorservice for handling admin products.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /products
func (s *Server) Initialize() {
	server := gin.Default()

	productsRoutes := server.Group("/api/rest_admin/products/")
	productsRoutes.GET("/", s.Handlers.Health)
	productsRoutes.POST("/", s.Handlers.AddProduct)
	productsRoutes.GET("/all", s.Handlers.GetProducts)
	productsRoutes.PUT("/:id", s.Handlers.UpdateProduct)
	productsRoutes.DELETE("/:id", s.Handlers.DeleteProduct)
	productsRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(config.Config["PORT"]))
}
