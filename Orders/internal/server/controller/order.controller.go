package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/api"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/literals"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/services"
)

type OrderController struct {
	service services.IService
}

func NewOrderController() *OrderController {
	return &OrderController{
		services.NewOrderService(),
	}
}

// Get Category By ID
// @Summary Get Single Category By Id
// @Tags Category
// @Param _id path string true "category_id"
// @Success 200 {object} models.Category
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /categories/{_id} [get]
func (cc *OrderController) GetByID(c *gin.Context) {
	id := c.GetString(literals.OrderIdKey)
	cat, err := cc.service.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (cc *OrderController) GetByStatus(c *gin.Context) {
	status := c.GetInt(literals.StatusKey)
	fmt.Println("status---", status)
	cat, err := cc.service.GetByStatus(status)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

func (cc *OrderController) GetByCustomer(c *gin.Context) {
	customerId := c.GetString(literals.CustomerIdKey)

	cat, err := cc.service.GetByCustomerId(customerId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

// Get All Categories
// @Summary Get All Categories
// @Tags Category
// @Success 200 {array} models.Category
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /categories [get]
func (cc *OrderController) GetAll(c *gin.Context) {
	res, err := cc.service.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create Category
// @Summary Create new Category, id is auto-assigned based on timestamp.
// @Tags Category
// @Param   category body models.Category true  "category"
// @Success 200 {object} models.Category
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /categories [post]
func (cc *OrderController) Create(c *gin.Context) {
	nc, _ := c.Get(literals.OrderBodyKey)
	newCat := nc.(models.Order)
	saved, err := cc.service.Create(newCat)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, saved)
}

// Update Category By ID
// @Summary Update Category By Id, Updated Category is returned
// @Tags Category
// @Param _id path string true "category_id"
// @Param   category body models.Category false  "category"
// @Success 200 {object} models.Category
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /categories/{_id} [put]
func (cc *OrderController) Update(c *gin.Context) {

	id := c.GetString(literals.OrderIdKey)
	nc, _ := c.Get(literals.OrderBodyKey)
	newCat := nc.(models.Order)

	saved, err := cc.service.UpdateByID(id, newCat)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, saved)
}

// Delete Category By ID
// @Summary Delete Single Category By Id
// @Tags Category
// @Param _id path string true "category_id"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /categories/{_id} [delete]
func (cc *OrderController) Delete(c *gin.Context) {
	id := c.GetString(literals.OrderIdKey)
	if err := cc.service.DeleteByID(id); err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.Status(http.StatusOK)
}

// Delete Multiple Categories
// @Summary Delete Multiple Categories
// @Tags Category
// @Param   category_ids body []int true  "category_ids"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErrors
// @Router /categories [delete]
//func (cc *CategoryController) DeleteMultiple(c *gin.Context) {
//	var ids []int
//	if err := c.BindJSON(&ids); err != nil {
//		c.AbortWithStatusJSON(api.ServerErr(err))
//		return
//	}
//
//	if errs := cc.service.DeleteMultiple(ids); errs != nil {
//		c.AbortWithStatusJSON(api.ServerErrs(errs))
//		return
//	}
//
//	c.Status(http.StatusOK)
//}
