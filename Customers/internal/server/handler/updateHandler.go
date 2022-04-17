package handler
import (
	"github.com/gin-gonic/gin"
	model "customers/internal/dao"
	db "customers/config"
	"customers/internal/server/service"
	"fmt"
)

// Update Customer godoc
// @Summary      Customer Update Route
// @Description  API to Update Customers-admin's Customer
// @Tags         Customer Update
// @Accept       json
// @Produce      json
// @Param        id  path string true "Customer Id"
// @Param        customer body {}  true "update details"
// @Success      200  {object} 	dao.Customer
// @Failure		 500   
// @Router       /customers/{id} [put]
func UpdateCustomer(ctx *gin.Context)(int,model.Customer,string){
	id_string:=ctx.Param("id");
	fmt.Println(id_string)
	var customer model.Customer
	ctx.BindJSON(&customer)
	db:=db.GetDB()
	return service.UpdateService(customer,id_string,db)
}
