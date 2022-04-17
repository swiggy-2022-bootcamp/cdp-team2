package handler
import (
	"github.com/gin-gonic/gin"
	model "customers/internal/dao"
	db "customers/config"
	"customers/internal/server/service"
)






// Get Customer by email godoc
// @Summary      Customer Get Route
// @Description  API to Get Customers-admin's Customer
// @Tags         Customer Get
// @Accept       json
// @Produce      json
// @Param        email path string true "Customer Email"
// @Success      200  {object} 	dao.Customer
// @Failure      500   
// @Router       /customers/email/{email} [get]
func GetCustomerEmail(ctx *gin.Context)(int,model.Customer,string){
	db:=db.GetDB()
	email:=ctx.Param("email");
	return service.GetService(email,db);
}
