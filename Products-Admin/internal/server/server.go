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

// @host localhost:8001
// @BasePath /api/rest_admin/products/
func (s *Server) Initialize() {
	server := gin.Default()

	baseRoutes := server.Group("/api/rest_admin")

	/*
	*	Product Specific Routes
	 */
	productsRoutes := baseRoutes.Group("/products/")
	productsRoutes.GET("/health", s.Handlers.Health)
	productsRoutes.POST("/", s.Handlers.AddProduct)
	productsRoutes.GET("/", s.Handlers.GetProducts)
	productsRoutes.PUT("/:id", s.Handlers.UpdateProduct)
	productsRoutes.DELETE("/:id", s.Handlers.DeleteProduct)

	/*
	* Search Specific Routes
	* Seach By
	* 	- Limit
	*	- Page
	* 	- Keyword
	*	- CategoryID
	* 	- ManufacturerID
	*	- Tag
	*	- date_added_from
	* 	- data_added_to
	*	- date_modified_from
	*	- date_modified_to
	 */
	searchRoutes := productsRoutes.Group("/search")
	searchRoutes.GET("/limit/:limit", nil)
	searchRoutes.GET("/page/:page", nil)
	searchRoutes.GET("/keyword/:keyword", s.Handlers.SearchByKeyword)
	searchRoutes.GET("/category/:id", s.Handlers.SearchByCategoryID)
	searchRoutes.GET("/manufacturer/:id", s.Handlers.SearchByManufacturerID)
	searchRoutes.GET("/tag/:tag", nil)
	searchRoutes.GET("/start/:price", s.Handlers.SearchByStartPrice)

	/*
	*	SwaggerUI Docs route
	 */
	productsRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(config.Config["PORT"]))
}
