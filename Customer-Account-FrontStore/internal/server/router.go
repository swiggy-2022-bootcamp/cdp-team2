package server
import (
	"github.com/gin-gonic/gin"
	"customer-account/internal/server/handler"
)


func CreateRoute(g *gin.RouterGroup){
	cont:=handler.NewCustomerController()
	g.POST("/",func(ctx *gin.Context){
		response,err:=cont.CreateAccount(ctx);
		if err==nil{
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,err.Error());
		}
	})
}



func CustomerRoute(g *gin.RouterGroup){

	cont:=handler.NewCustomerController()

	g.PUT("/:id",func(ctx *gin.Context){
		response,err:=cont.UpdateAccount(ctx);
		if err==nil{
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,err.Error());
		}
	})

	g.GET("/:id",func(ctx *gin.Context){
		response,err:=cont.ReadAccount(ctx);
 		if err==nil{
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,err.Error());
		}
	})
  
}


func HealthRoute(g *gin.RouterGroup){
	g.GET("/",func(ctx *gin.Context){
		status,message:=handler.Health(ctx)
		ctx.JSON(status,gin.H{"message":message});
	})
	
}