package handler
import (
	"github.com/gin-gonic/gin"
	model "customers/internal/dao"
	"customers/internal/server/service"
	db "customers/config"
	"customers/util"
)




// Create Customer godoc
// @Summary      Customer Create Route
// @Description  API to Create Customers-admin's Customer
// @Tags         Customer Create
// @Accept       json
// @Produce      json
// @Param        customer body dao.Customer ture "Customer Body"
// @Success      200  {object} 	dao.Customer
// @Failure      500  
// @Router       /customers [post]
func CreateCustomer(ctx *gin.Context)(int,model.Customer,string){
	var customer model.Customer
	ctx.BindJSON(&customer)
	customer.Id=util.Generate()
	db:=db.GetDB()
	return service.CreateService(customer,db)
}

