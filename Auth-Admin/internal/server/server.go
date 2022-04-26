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

// @title Swagger Admin Auth Microservice
// @version 1.0
// @description Micorservice for handling Admin Auth.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 35.84.28.237:8000
// @BasePath /auth
func (s *Server) Initialize() {
	server := gin.Default()

	adminAuthRoutes := server.Group("/auth/")
	adminAuthRoutes.POST("/login", s.Handlers.Login)
	adminAuthRoutes.POST("/oAuth", s.Middlewares.CheckBasicAuthMiddleware, s.Middlewares.CheckAdminRole, s.Handlers.OAuth)
	adminAuthRoutes.GET("/logout", s.Middlewares.CheckAuthMiddleware, s.Handlers.Logout)
	adminAuthRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	adminAuthRoutes.GET("/", s.Handlers.Health)

	log.Fatal(server.Run(config.Config["PORT"]))
}
