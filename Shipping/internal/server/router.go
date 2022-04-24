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

	shippingGrp := r.Group("/shipping-address")
	{
		shippingGrp.POST("/", controller.BindOrder, cont.Create)
		shippingGrp.GET("/customer/:customerId", controller.BindCustomer, cont.GetByCustomer)
		shippingGrp.POST("/set/address", controller.BindAddressOrder, cont.SetAddressToOrder)
	}

	return r
}
