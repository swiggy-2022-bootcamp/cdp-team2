package server
import (
	"github.com/gin-gonic/gin"
	"customers/internal/server/handler"
)

func CustomerRoute(g *gin.RouterGroup){
	cont:=handler.NewCustomerController()
	g.POST("/",func(ctx *gin.Context){
		response,err:=cont.CreateCustomer(ctx);
 		if err==nil{
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,gin.H{"message":err.Error()});
		}
	})

	g.PUT("/:id",func(ctx *gin.Context){
		response,err:=cont.UpdateCustomer(ctx);
		if err==nil{
			
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,gin.H{"message":err.Error()});
		}
	})

	g.GET("/:id",func(ctx *gin.Context){
		response,err:=cont.ReadCustomer(ctx);
		if err==nil{
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,gin.H{"message":err.Error()});
		}
	})

	g.DELETE("/:id",func(ctx *gin.Context){
		response,_:=cont.DeleteCustomer(ctx);
		ctx.JSON(200,gin.H{"message":response});

	})

	g.GET("/email/:email",func(ctx *gin.Context){
		response,err:=cont.ReadCustomerByEmail(ctx);
		if err==nil{
			ctx.JSON(200,response);
		} else{
			ctx.JSON(400,gin.H{"message":err.Error()});
		}
	})

}

// Health godoc
// @Summary      Health Check Route
// @Description  API to check Customer-Admin's health
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object} 	string
// @Router       /health [get]
func HealthRoute(g *gin.RouterGroup){
	g.GET("/",func(ctx *gin.Context){
		// status,message:=handler.Health(ctx)
		ctx.JSON(200,gin.H{"message":"server up"});
	})
}