package controller

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/literals"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, literals.Health_OK)
}
