package handler

import (
	"github.com/gin-gonic/gin"
	"customer-account/dao/models"
 	services "customer-account/internal/server/service"
)

type CustomerController struct {
	service services.IService
}

//go:generate mockgen --destination=../mocks/customerMock.go  IService
func NewCustomerController() *CustomerController {
	return &CustomerController{
		services.NewCustomerService(),
	}
}


type ICustomer interface {
	CreateService(ctx *gin.Context) (models.Account, error)
	ReadService(ctx *gin.Context)( models.Account2,error)
	UpdateService(ctx *gin.Context)(models.Account,error)
}





