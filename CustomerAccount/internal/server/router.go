package server
import (
	
	"github.com/gin-gonic/gin"
	"customer-account/internal/server/handler"
)

func CustomerRoute(g *gin.RouterGroup){
	g.GET("/",func(ctx *gin.Context){
		
	})
}


func HealthRoute(g *gin.RouterGroup){
	g.GET("/",func(ctx *gin.Context){
		status,message:=handler.Health(ctx)
		ctx.JSON(status,gin.H{"message":message});
	})
}