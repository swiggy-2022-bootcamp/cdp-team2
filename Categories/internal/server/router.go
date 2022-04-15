package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/docs"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/server/controllers"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//health Check
	r.GET("/health", controllers.Health)

	cont := controllers.NewCategoryController()

	catGrp := r.Group("/categories")
	{
		catGrp.GET("/", cont.GetAll)
		catGrp.GET("/:_id", controllers.BindId, cont.GetByID)
		catGrp.POST("/", controllers.BindCategory, cont.Create)
		catGrp.PUT("/:_id", controllers.BindId, controllers.BindCategory, cont.Update)
		catGrp.DELETE("/:_id", controllers.BindId, cont.Delete)
		catGrp.DELETE("/", cont.DeleteMultiple)
	}

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
