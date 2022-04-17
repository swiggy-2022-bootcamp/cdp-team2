package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/api"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/literals"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/services"
)

type ShippingAddressController struct {
	service services.IService
}

func NewOShippingAddressController() *ShippingAddressController {
	return &ShippingAddressController{
		services.NewShippingAddressService(),
	}
}

func (cc *ShippingAddressController) GetByCustomer(c *gin.Context) {
	customerId := c.GetInt(literals.CustomerIdKey)

	cat, err := cc.service.GetByCustomerId(customerId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (cc *ShippingAddressController) Create(c *gin.Context) {
	nc, _ := c.Get(literals.AddressBodyKey)
	newCat := nc.(models.ShippingAddress)
	saved, err := cc.service.Create(newCat)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, saved)
}
