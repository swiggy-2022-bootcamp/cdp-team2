package handler

import (
	"github.com/gin-gonic/gin"
	model "customer-account/internal/dao"
	db "customer-account/config"
	service "customer-account/internal/server/service"
)
// Get Customer Account by Id godoc
// @Summary      Customer Account Get Route
// @Description  API to Get Customers-Frontstore's Customer
// @Tags         Customer Account Get
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Customer Id"
// @Success      200  {object} 	dao.Account2
// @Failure 	 500  {object}  string 
// @Router       /account/{id} [get]
func ReadAccount(ctx *gin.Context)(int,model.Account2,string){
	db:=db.GetDB()
	id_string:=ctx.Param("id");
	return service.ReadService(id_string,db)
}
