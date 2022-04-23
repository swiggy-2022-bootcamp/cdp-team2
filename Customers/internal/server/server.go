package server

import (
 
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"  //gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"	//swagger embed files
	_ "customers/docs"
	// model "customers/config"
	"customers/config"
)


// @title Customers API
// @version 1.0
// @description This is customers CRUD service.
// @termsOfService demo.com

// @contact.name API Support
// @contact.url http://demo.com/support

// @host localhost:8091
// @BasePath /

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func RunServer()error{
	config.FromEnv()
 	server:=gin.Default()

	customerRoute:=server.Group("/customers")
	{
		CustomerRoute(customerRoute.Group("/"))
	}

	healthRoute:=server.Group("/health")
	{
		HealthRoute(healthRoute.Group("/"))
	}
	 
	server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	go GrpcServer()
	err:=server.Run(":"+config.Server["PORT"])
	if err!=nil{
		return err
	}
	return nil;
}