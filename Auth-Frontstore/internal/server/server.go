package server

import (
	"log"

	"github.com/auth-frontstore-service/config"
	_ "github.com/auth-frontstore-service/docs"
	"github.com/auth-frontstore-service/internal/core/ports"
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

// @title Swagger Frontstore Auth Microservice
// @version 1.0
// @description Micorservice for handling Frontstore Auth.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8088
// @BasePath /auth
func (s *Server) Initialize() {
	server := gin.Default()

	adminAuthRoutes := server.Group("/auth/")
	adminAuthRoutes.POST("/login", s.Handlers.Login)
	adminAuthRoutes.POST("/oAuth/token", s.Middlewares.CheckBasicAuthMiddleware, s.Handlers.OAuth)
	adminAuthRoutes.GET("/logout", s.Middlewares.CheckAuthMiddleware, s.Handlers.Logout)
	adminAuthRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	adminAuthRoutes.GET("/", s.Handlers.Health)

	log.Fatal(server.Run(config.Config["PORT"]))
}
