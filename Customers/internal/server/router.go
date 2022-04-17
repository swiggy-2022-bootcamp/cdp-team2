package server
import (
	
	"github.com/gin-gonic/gin"
	"customers/internal/server/handler"
)

func CustomerRoute(g *gin.RouterGroup){
	g.POST("/",func(ctx *gin.Context){
		status,response,error:=handler.CreateCustomer(ctx);
		if status==200{
			ctx.JSON(status,response);
		} else{
			ctx.JSON(status,error);
		}
	})

	g.PUT("/:id",func(ctx *gin.Context){
		status,response,error:=handler.UpdateCustomer(ctx);
		if status==200{
			ctx.JSON(status,response);
		} else{
			ctx.JSON(status,error);
		}
	})

	g.GET("/:id",func(ctx *gin.Context){
		status,response,error:=handler.ReadCustomer(ctx);
		if status==200{
			ctx.JSON(status,response);
		} else{
			ctx.JSON(status,error);
		}
	})

	g.DELETE("/:id",func(ctx *gin.Context){
		status,response:=handler.DeleteCustomer(ctx);
		ctx.JSON(status,response);

	})

	g.GET("/email/:email",func(ctx *gin.Context){
		status,response,error:=handler.GetCustomerEmail(ctx);
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