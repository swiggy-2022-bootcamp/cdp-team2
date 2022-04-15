package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/docs"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/server/controllers"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(gin.Recovery())

	//health Check
	r.GET("/health", controllers.Health)

	cont := controllers.NewCheckoutController()

	r.POST("/checkout/:cart_id", cont.StartCheckout)
	r.PUT("/rewards", cont.ApplyReward)
	r.POST("/pay", cont.Pay)
	r.GET("/order/:order_id", cont.GetOrder)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
