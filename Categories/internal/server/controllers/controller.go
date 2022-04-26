package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/api"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/literals"
)

func BindId(c *gin.Context) {
	idstr := c.Param(literals.CatIdKey)
	if idstr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CatIDNotFound})
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CatIDNotFound})
		return
	}
	c.Set(literals.CatIdKey, id)
	c.Next()
}

func BindCategory(c *gin.Context) {
	cat := models.Category{}
	if err := c.BindJSON(&cat); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CatBodyKey})
		return
	}
	c.Set(literals.CatBodyKey, cat)
	c.Next()
}
