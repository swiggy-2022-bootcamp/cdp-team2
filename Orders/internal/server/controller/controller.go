package controller

import (
	"fmt"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/literals"
)

func BindId(c *gin.Context) {
	id := c.Param(literals.OrderIdKey)
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.OrderNotFound})
		return
	}

	c.Set(literals.OrderIdKey, id)
	c.Next()
}

func BindStatus(c *gin.Context) {
	status, _ := strconv.Atoi(c.Param(literals.StatusKey))
	if status == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.OrderNotFound})
		return
	}

	if status > 3 || status < 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.StatusNotValid})
		return

	}

	c.Set(literals.StatusKey, status)
	c.Next()
}

func BindCustomer(c *gin.Context) {
	id := c.Param(literals.CustomerIdKey)
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.OrderNotFound})
		return
	}

	c.Set(literals.CustomerIdKey, id)
	c.Next()
}

func BindOrder(c *gin.Context) {
	cat := models.Order{}
	if err := c.BindJSON(&cat); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.OrderBodyKey})
		return
	}

	//if cat.Status > 3 || cat.Status <= 0 {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.StatusNotValid})
	//	return
	//}
	c.Set(literals.OrderBodyKey, cat)
	c.Next()
}
