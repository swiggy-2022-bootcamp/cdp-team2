package handler
import (
	"github.com/gin-gonic/gin"
)
// Health godoc
// @Summary      Health Check Route
// @Description  API to check Customers-admin's health
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object} 	string
// @Router       /health [get]
func Health(ctx *gin.Context)(int,string){
	return 200,"Server_Up"
}