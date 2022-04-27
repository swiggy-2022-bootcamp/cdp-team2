package handler

import (
	"github.com/gin-gonic/gin"
	"customer-account/dao/models"
	
)
// Get Customer Account by Id godoc
// @Summary      Customer Account Get Route
// @Description  API to Get Customers-Frontstore's Customer
// @Tags         Customer Account Get
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Customer Id"
// @Success      200  {object} 	models.Account2
// @Failure 	 500  {object}  string 
// @Router       /account/{id} [get]
func (cc *CustomerController)ReadAccount(ctx *gin.Context)(models.Account2,error){
	id_string:=ctx.Param("id");
	return cc.service.ReadService(id_string)
}

// Create godoc
// @Summary      Account Create Route
// @Description  API to Create Customers-Frontstore Customer
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        account body models.Account true "Account info"
// @Success      200  {object} 	models.Account
// @Failure 	 500  {object}	string
// @Router       /register [post]
func (cc *CustomerController)CreateAccount(ctx *gin.Context)(models.Account,error){
	var account models.Account
 	ctx.BindJSON(&account)
	return cc.service.CreateService(account)
}



// Update Customer Account godoc
// @Summary      Customer Account Create Route
// @Description  API to Update Customer-Frontstore's Customer
// @Tags         Customer Account Update
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer Id"
// @Param        account body {} true "Account Update"
// @Success      200  {object} 	models.Account
// @Failure      500  {object}  string
// @Router       /account/{id} [put]
func (cc *CustomerController)UpdateAccount(ctx *gin.Context)(models.Account,error){
	id_string:=ctx.Param("id");
 	var customer models.Account
	ctx.BindJSON(&customer)
	return cc.service.UpdateService(id_string,customer)
}
