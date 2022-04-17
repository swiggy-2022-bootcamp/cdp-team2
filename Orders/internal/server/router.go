package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/server/controller"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	//health Check
	r.GET("/health", controller.Health)

	cont := controller.NewOrderController()

	catGrp := r.Group("/orders")
	{
		catGrp.GET("/", cont.GetAll)
		catGrp.GET("/:id", controller.BindId, cont.GetByID)
		catGrp.POST("/", controller.BindOrder, cont.Create)
		catGrp.GET("/status/:status", controller.BindStatus, cont.GetByStatus)
		catGrp.PUT("/:id", controller.BindId, controller.BindOrder, cont.Update)
		catGrp.DELETE("/:id", controller.BindId, cont.Delete)
		catGrp.GET("/customer/:customerId", controller.BindCustomer, cont.GetByCustomer)
	}

	return r
}
