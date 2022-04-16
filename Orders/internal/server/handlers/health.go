package handlers

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/literals"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, literals.Health_OK)
}
