package server

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"  //gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"	//swagger embed files
	_ "customer-account/docs"
	"customer-account/config"
	"fmt"
)



// @title Account API
// @version 1.0
// @description This is Account CRUD service.
// @termsOfService demo.com

// @contact.name API Support
 
// @host localhost:config.Server["PORT"]
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func RunServer(){
	fmt.Println("Server:Config",config.Server,config.AWS)
	server:=gin.Default()
	createRoute:=server.Group("/register")
	{
		CreateRoute(createRoute.Group("/"))
	}

	customerRoute:=server.Group("/account")
	{
		CustomerRoute(customerRoute.Group("/"))
	}

	healthRoute:=server.Group("/health")
	{
		HealthRoute(healthRoute.Group("/"))
	}

	server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	go StartGrpc()
	server.Run(":"+config.Server["PORT"])
}