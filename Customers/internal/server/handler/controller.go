package handler

import (
	"github.com/gin-gonic/gin"
	"customers/dao/models"
 	services "customers/internal/server/service"
)

type CustomerController struct {
	service services.IService
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		services.NewCustomerService(),
	}
}


type ICustomer interface {
	CreateService(ctx *gin.Context) (models.Customer, error)
	ReadService(ctx *gin.Context)( models.Customer,error)
	UpdateService(ctx *gin.Context)(models.Customer,error)
}





