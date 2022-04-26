package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/internal/literals"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, literals.Health_OK)
}
