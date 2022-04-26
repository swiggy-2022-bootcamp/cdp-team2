package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/literals"
)

// Health Check
// @Summary Health Check
// @Tags Health
// @Success 200 {string} string "ok"
// @Failure 500
// @Router /health [get]
func Health(c *gin.Context) {
	c.String(http.StatusOK, literals.Health_OK)
}
