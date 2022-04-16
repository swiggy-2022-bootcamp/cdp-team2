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
	idstr := c.Param(literals.OrderIdKey)
	if idstr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CatIDNotFound})
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CatIDNotFound})
		return
	}
	c.Set(literals.OrderIdKey, id)
	c.Next()
}

func BindStatus(c *gin.Context) {
	status := c.Param(literals.StatusKey)
	fmt.Println(status)
	if status == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CatIDNotFound})
		return
	}

	c.Set(literals.StatusKey, status)
	c.Next()
}

func BindOrder(c *gin.Context) {
	cat := models.Order{}
	if err := c.BindJSON(&cat); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.OrderBodyKey})
		return
	}
	c.Set(literals.OrderBodyKey, cat)
	c.Next()
}
