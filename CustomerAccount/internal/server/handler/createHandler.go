package handler

import (
	"github.com/gin-gonic/gin"
	model "customer-account/internal/dao"
	db "customer-account/config"
	service "customer-account/internal/server/service"
	util "customer-account/util"
)

// Create godoc
// @Summary      Account Create Route
// @Description  API to Create Customers-Frontstore Customer
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param        account body dao.Account true "Account info"
// @Success      200  {object} 	dao.Account
// @Failure 	 500  {object}	string
// @Router       /register [post]
func CreateAccount(ctx *gin.Context)(int,model.Account,string){
	var account model.Account
 	ctx.BindJSON(&account)
 	account.Id=util.Generate()
	db:=db.GetDB()
	return service.CreateService(account,db)
}