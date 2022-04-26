package controller

import (
	"fmt"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/literals"
)

func BindCustomer(c *gin.Context) {
	idstr := c.Param(literals.CustomerIdKey)
	if idstr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.AddressNotFound})
		return
	}
	
	c.Set(literals.CustomerIdKey, idstr)
	c.Next()
}

func BindOrder(c *gin.Context) {
	shippingAddress := models.ShippingAddress{}
	if err := c.BindJSON(&shippingAddress); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.AddressBodyKey})
		return
	}
	c.Set(literals.AddressBodyKey, shippingAddress)
	c.Next()
}

func BindAddressOrder(c *gin.Context) {
	shippingAddress := api.SetAddressToOrderRequest{}
	if err := c.BindJSON(&shippingAddress); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.AddressBodyKey})
		return
	}
	c.Set(literals.OrderBodyKey, shippingAddress.OrderID)
	c.Set(literals.AddressBodyKey, shippingAddress.AddressId)
	c.Next()
}
