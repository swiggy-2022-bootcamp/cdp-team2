package server

import (
 
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"  //gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"	//swagger embed files
	_ "customers/docs"

)



// @title Customers API
// @version 1.0
// @description This is customers CRUD service.
// @termsOfService demo.com

// @contact.name API Support
// @contact.url http://demo.com/support

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func RunServer(){
 
	server:=gin.Default()

	customerRoute:=server.Group("/customer")
	{
		CustomerRoute(customerRoute.Group("/"))
	}

	healthRoute:=server.Group("/health")
	{
		HealthRoute(healthRoute.Group("/"))
	}
	 
	server.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":8080")
}