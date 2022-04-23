package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/api"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/literals"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/services"
)

type CheckoutController struct {
	service services.ICheckoutService
}

func NewCheckoutController() (*CheckoutController, error) {
	s, err := services.NewCheckoutService()
	if err != nil {
		return nil, err
	}
	return &CheckoutController{s}, nil
}

// Start Checkout
// @Summary Start Checkout.
// @Tags Checkout
// @Param cart_id path string true "cart_id"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /checkout/{cart_id} [post]
func (cc *CheckoutController) StartCheckout(c *gin.Context) {

	idstr := c.Param(literals.CartIdKey)
	if idstr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.CartIdNotFound})
		return
	}
	cartId, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(api.UserErr(err))
		return
	}

	log.Printf("Starting Checkout for cart : %d", cartId)

	err = cc.service.StartCheckout(cartId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// Apply Reward
// @Summary Apply Reward on Order.
// @Tags Checkout
// @Param   applyReward body api.ApplyRewardRequest true "apply reward input"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /rewards [put]
func (cc *CheckoutController) ApplyReward(c *gin.Context) {

	req := api.ApplyRewardRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(api.UserErr(err))
		return
	}

	err := cc.service.ApplyReward(req.OrderID, req.Rewards)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// Pay
// @Summary Pay Order price.
// @Tags Checkout
// @Param   pay body api.PayRequest true "pay & confirm order input"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /pay [post]
func (cc *CheckoutController) Pay(c *gin.Context) {

	req := api.PayRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(api.UserErr(err))
		return
	}

	err := cc.service.ConfirmOrder(req.OrderID)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// Get Order
// @Summary Get Order .
// @Tags Checkout
// @Param order_id path string true "order_id"
// @Success 200
// @Failure 500 {object} api.ApiResponseWithErr
// @Router /order/{order_id} [get]
func (cc *CheckoutController) GetOrder(c *gin.Context) {

	idstr := c.Param(literals.OrderIdKey)
	if idstr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.ApiResponseWithErr{literals.OrderIdNotFound})
		return
	}
	orderId, err := strconv.Atoi(idstr)
	if err != nil {
		c.AbortWithStatusJSON(api.UserErr(err))
		return
	}

	err = cc.service.GetOrder(orderId)
	if err != nil {
		c.AbortWithStatusJSON(api.ServerErr(err))
		return
	}

	c.Status(http.StatusOK)
}
