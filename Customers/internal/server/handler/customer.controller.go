package handler

import (
	"github.com/gin-gonic/gin"
	"customers/dao/models"
	
)
// Get Customer by Id godoc
// @Summary      Customer Get Route
// @Description  API to Get Customer-Admin's Customer
// @Tags         Customer Get
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Customer Id"
// @Success      200  {object} 	models.Customer
// @Failure 	 500  {object}  string 
// @Router       /customers/{id} [get]
func (cc *CustomerController)ReadCustomer(ctx *gin.Context)(models.Customer,error){
	id_string:=ctx.Param("id");
	return cc.service.ReadService(id_string)
}

// Get Customer by Id godoc
// @Summary      Customer Get By Email Route
// @Description  API to Get Customer-Admin's Customer
// @Tags         Customer Get
// @Accept       json
// @Produce      json
// @Param 		 email path string true "Customer email"
// @Success      200  {object} 	models.Customer
// @Failure 	 500  {object}  string 
// @Router       /customers/email/{email} [get]
func (cc *CustomerController)ReadCustomerByEmail(ctx *gin.Context)(models.Customer,error){
	id_string:=ctx.Param("email");
	return cc.service.ReadByEmailService(id_string)
}

// Delete godoc
// @Summary      Customer Delete Route
// @Description  API to Delete Customer-Admin's Customer
// @Tags         Delete
// @Accept       json
// @Produce      json
// @Param 		 id path string true "Customer Id"
// @Success      200  {object} 	models.Customer
// @Failure 	 500  {object}	string
// @Router       /customers/{id} [delete]
func (cc *CustomerController)DeleteCustomer(ctx *gin.Context)(bool,error){
	id_string:=ctx.Param("id");
	return cc.service.DeleteService(id_string)
}
// Create godoc
// @Summary      Customer Create Route
// @Description  API to Create Customer-Admin's Customer
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        Customer body models.Customer true "Customer info"
// @Success      200  {object} 	models.Customer
// @Failure 	 500  {object}	string
// @Router       /customers [post]
func (cc *CustomerController)CreateCustomer(ctx *gin.Context)(models.Customer,error){
	var account models.Customer
 	ctx.BindJSON(&account)
	return cc.service.CreateService(account)
}



// Update Customer godoc
// @Summary      Customer Update Route
// @Description  API to Update Customer-Admin's Customer
// @Tags         Customer Update
// @Accept       json
// @Produce      json
// @Param        id path string true "Customer Id"
// @Param        customer body {} true "Customer Update"
// @Success      200  {object} 	models.Customer
// @Failure      500  {object}  string
// @Router       /customers/{id} [put]
func (cc *CustomerController)UpdateCustomer(ctx *gin.Context)(models.Customer,error){
	id_string:=ctx.Param("id");
 	var customer models.Customer
	ctx.BindJSON(&customer)
	return cc.service.UpdateService(id_string,customer)
}
