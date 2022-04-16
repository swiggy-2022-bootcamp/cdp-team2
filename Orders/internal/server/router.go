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
		//catGrp.GET("/", controller.GetAll)
		catGrp.GET("/:id", controller.BindId, cont.GetByID)
		catGrp.POST("/", controller.BindOrder, cont.Create)
		//catGrp.PUT("/:_id", controller.BindId, controller.BindCategory, cont.Update)
		//catGrp.DELETE("/:_id", controller.BindId, cont.Delete)
		//catGrp.DELETE("/", controller.DeleteMultiple)
	}

	return r
}
