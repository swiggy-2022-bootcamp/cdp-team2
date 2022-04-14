package server

import (
	"log"

	"github.com/auth-admin-service/config"
	_ "github.com/auth-admin-service/docs"
	"github.com/auth-admin-service/internal/core/ports"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Handlers    ports.IHandlers
	Middlewares ports.IMiddlewares
}

func NewServer(handlers ports.IHandlers, middlewares ports.IMiddlewares) *Server {
	return &Server{
		Handlers:    handlers,
		Middlewares: middlewares,
	}
}

// @title Swagger adminAuth Admin Microservice
// @version 1.0
// @description Micorservice for handling admin adminAuth.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /adminAuth
func (s *Server) Initialize() {
	server := gin.Default()

	adminAuthRoutes := server.Group("/auth/")
	adminAuthRoutes.POST("/login", s.Handlers.Login)
	adminAuthRoutes.POST("/oAuth", s.Middlewares.CheckAuthMiddleware, s.Middlewares.CheckAdminRole, s.Handlers.OAuth)
	adminAuthRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	adminAuthRoutes.GET("/", s.Handlers.Health)
	adminAuthRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(server.Run(config.Config["PORT"]))
}
