package handler

import (
	"github.com/gin-gonic/gin"
	model "customer-account/internal/dao"
	db "customer-account/config"
	service "customer-account/internal/server/service"
)

// Update Customer Account godoc
// @Summary      Customer Account Create Route
// @Description  API to Update Customer-Frontstore's Customer
// @Tags         Customer Account Update
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer Id"
// @Param        account body {} true "Account Update"
// @Success      200  {object} 	dao.Account
// @Failure      500  {object}  string
// @Router       /account/{id} [put]
func UpdateAccount(ctx *gin.Context)(int,model.Account,string){
	id_string:=ctx.Param("id");
 	var customer model.Account
	ctx.BindJSON(&customer)
	db:=db.GetDB()
	return service.UpdateService(id_string,customer,db)
}
