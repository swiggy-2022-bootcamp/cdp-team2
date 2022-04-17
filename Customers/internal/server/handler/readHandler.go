package handler
import (
	"github.com/gin-gonic/gin"
	model "customers/internal/dao"
	db "customers/config"
	"customers/internal/server/service"
)

// Get Customer by Id godoc
// @Summary      Customer Get Route
// @Description  API to Get Customers-admin's Customer
// @Tags         Customer Get
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer Id"
// @Success      200  {object} 	dao.Customer
// @Failure 	 500  
// @Router       /customers/{id} [get]
func ReadCustomer(ctx *gin.Context)(int,model.Customer,string){
	db:=db.GetDB()
	id_string:=ctx.Param("id");
	return service.ReadService(id_string,db)
}