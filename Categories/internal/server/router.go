package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/docs"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/server/handlers"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	//health Check
	r.GET("/health", handlers.Health)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
