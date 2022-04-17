package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/server/controller"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	//health Check
	r.GET("/health", controller.Health)

	cont := controller.NewOShippingAddressController()

	catGrp := r.Group("/shipping-address")
	{
		catGrp.POST("/", controller.BindOrder, cont.Create)
		catGrp.GET("/customer/:customerId", controller.BindCustomer, cont.GetByCustomer)
	}

	return r
}
