package server
import (
	
	"github.com/gin-gonic/gin"
	"customer-account/internal/server/handler"
)

func CreateRoute(g *gin.RouterGroup){
	g.POST("/",func(ctx *gin.Context){
		status,response,error:=handler.CreateAccount(ctx);
		if status==200{
			ctx.JSON(status,response);
		} else{
			ctx.JSON(status,error);
		}
	})
}

func CustomerRoute(g *gin.RouterGroup){
	

	g.PUT("/:id",func(ctx *gin.Context){
		status,response,error:=handler.UpdateAccount(ctx);
		if status==200{
			ctx.JSON(status,response);
		} else{
			ctx.JSON(status,error);
		}
	})

	g.GET("/:id",func(ctx *gin.Context){
		status,response,error:=handler.ReadAccount(ctx);
		if status==200{
			ctx.JSON(status,response);
		} else{
			ctx.JSON(status,error);
		}
	})
  
}


func HealthRoute(g *gin.RouterGroup){
	g.GET("/",func(ctx *gin.Context){
		status,message:=handler.Health(ctx)
		ctx.JSON(status,gin.H{"message":message});
	})
}