package controller

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/services"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	_ "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/docs"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/api"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/literals"
)

type ShippingAddressController struct {
	service services.IService
}

func NewOShippingAddressController() *ShippingAddressController {
	return &ShippingAddressController{
		services.NewShippingAddressService(),
	}
}

// @Summary get shipping address of a customer
// @ID GetByCustomer
// @Param customerId path string true "customerId"
// @Produce json
// @Success 200
// @Router /shipping-address/customer/{customerId} [get]
func (cc *ShippingAddressController) GetByCustomer(c *gin.Context) {
	customerId := c.GetInt(literals.CustomerIdKey)

	shippingAddress, err := cc.service.GetByCustomerId(customerId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, shippingAddress)
}

// @Summary add shipping address to a customer
// @ID Create
// @Param data body object true "address"
// @Produce json
// @Success 200
// @Router /shipping-address/ [post]
func (cc *ShippingAddressController) Create(c *gin.Context) {
	nc, _ := c.Get(literals.AddressBodyKey)
	newshippingAddress := nc.(models.ShippingAddress)
	saved, err := cc.service.Create(newshippingAddress)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, saved)
}

// @Summary set shipping address to a order
// @ID SetAddressToOrder
// @Param orderId path string true "orderId"
// @Param addressId path string true "addressId"
// @Produce json
// @Success 200
// @Router /shipping-address/set/address [post]
func (cc *ShippingAddressController) SetAddressToOrder(c *gin.Context) {
	orderId := c.GetString(literals.OrderBodyKey)
	addressId := c.GetInt(literals.AddressBodyKey)

	res, err := cc.service.SetAddressToOrder(orderId, addressId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, res)
}
