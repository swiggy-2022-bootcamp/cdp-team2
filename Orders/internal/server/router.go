package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/server/handlers"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	//health Check
	r.GET("/health", handlers.Health)

	return r
}
