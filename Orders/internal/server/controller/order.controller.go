package controller

import (
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

// Get Order By ID
// @Summary Get a Order By Id
// @Tags Order
// @Param orderId path string true "orderId"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders/{orderId} [get]
func (cc *OrderController) GetByID(c *gin.Context) {
	id := c.GetString(literals.OrderIdKey)
	cat, err := cc.service.GetByID(id)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

// Get Order By status
// @Summary Get a Order By Status
// @Tags Order
// @Param statusId path int true "statusId"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders/status/{statusId} [get]
func (cc *OrderController) GetByStatus(c *gin.Context) {
	status := c.GetInt(literals.StatusKey)
	cat, err := cc.service.GetByStatus(status)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

// Get Order By status
// @Summary Get a Order By CustomerId
// @Tags Order
// @Param customerId path int true "customerId"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders/customer/{customerId} [get]
func (cc *OrderController) GetByCustomer(c *gin.Context) {
	customerId := c.GetString(literals.CustomerIdKey)

	cat, err := cc.service.GetByCustomerId(customerId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, cat)
}

// Get All Orders
// @Summary Get All Orders
// @Tags Order
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders [get]
func (cc *OrderController) GetAll(c *gin.Context) {
	res, err := cc.service.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create Order
// @Summary Create new Order, id is auto-assigned uuid.
// @Tags Order
// @Param  order body object true  "order"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders [post]
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

// Update Order By ID
// @Summary Update Order By Id
// @Tags Order
// @Param orderId path string true "orderId"
// @Param   order body  object true  "order"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders/{orderId} [put]
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

// Delete Order By ID
// @Summary Delete Single Order By Id
// @Tags Order
// @Param orderId path string true "orderId"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /orders/{orderId} [delete]
func (cc *OrderController) Delete(c *gin.Context) {
	id := c.GetString(literals.OrderIdKey)
	if err := cc.service.DeleteByID(id); err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}
	c.Status(http.StatusOK)
}
