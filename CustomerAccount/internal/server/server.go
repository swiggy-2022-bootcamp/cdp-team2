package server

import (
	"customer-account/config"
	_ "customer-account/docs"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   //gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" //swagger embed files
)

// @title Account API
// @version 1.0
// @description This is Account CRUD service.
// @termsOfService demo.com

// @contact.name API Support

// @host 35.84.28.237:30214
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func RunServer() {
	fmt.Println("Server:Config", config.Server, config.AWS)
	server := gin.Default()
	server.Use(cors.Default())
	createRoute := server.Group("/register")
	{
		CreateRoute(createRoute.Group("/"))
	}

	customerRoute := server.Group("/account")
	{
		CustomerRoute(customerRoute.Group("/"))
	}

	healthRoute := server.Group("/health")
	{
		HealthRoute(healthRoute.Group("/"))
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	go StartGrpc()
	server.Run(":" + config.Server["PORT"])
}
