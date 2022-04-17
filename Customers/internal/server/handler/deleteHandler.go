package handler
import (
	"github.com/gin-gonic/gin"
	db "customers/config"	
	"customers/internal/server/service"
)




// Delete Customer godoc
// @Summary      Customer Delete Route
// @Description  API to Delete Customers-admin's Customer
// @Tags         Customer Delete
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer Id"
// @Success      200  {object} 	string
// @Failure      500   
// @Router       /customers/{id} [delete]
func DeleteCustomer(ctx *gin.Context)(int,string){
	db:=db.GetDB()
	id_string:=ctx.Param("id");
	return service.DeleteService(id_string,db)
}
 

