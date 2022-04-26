package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/server/controller"
	middlewares "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/server/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())

	//health Check
	r.GET("/health", controller.Health)

	cont := controller.NewOShippingAddressController()

	shippingGrp := r.Group("/shipping-address")
	{
		shippingGrp.POST("/", middlewares.CheckAuthMiddleware, controller.BindOrder, cont.Create)
		shippingGrp.GET("/customer/:customerId", middlewares.CheckAuthMiddleware, controller.BindCustomer, cont.GetByCustomer)
		shippingGrp.POST("/set/address", middlewares.CheckAuthMiddleware, controller.BindAddressOrder, cont.SetAddressToOrder)
	}

	return r
}
