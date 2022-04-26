package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/server/controller"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())

	//health Check
	r.GET("/health", controller.Health)

	cont := controller.NewOrderController()

	orderGrp := r.Group("/orders")
	{
		orderGrp.GET("/", cont.GetAll)
		orderGrp.GET("/:id", controller.BindId, cont.GetByID)
		orderGrp.POST("/", controller.BindOrder, cont.Create)
		orderGrp.GET("/status/:status", controller.BindStatus, cont.GetByStatus)
		orderGrp.PUT("/:id", controller.BindId, controller.BindOrder, cont.Update)
		orderGrp.DELETE("/:id", controller.BindId, cont.Delete)
		orderGrp.GET("/customer/:customerId", controller.BindCustomer, cont.GetByCustomer)
	}

	return r
}
