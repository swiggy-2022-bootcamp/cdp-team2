package middlewares

import (
	"net/http"
	"strings"

	repo "github.com/auth-admin-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type MiddlewaresImpl struct {
	// TODO: Need to inject product services
}

func NewMiddlewares() *MiddlewaresImpl {
	return &MiddlewaresImpl{}
}

// Health godoc
// @Summary      Health Check Route
// @Description  API to check products-admin health
// @Tags         Health
// @Accept       json
// @Produce      plain/text
// @Success      200  {object} 	string
// @Router       / [get]

func (h *MiddlewaresImpl) CheckAuthMiddleware(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) > 0 {
		ok, err := repo.JWTManager.Verify(strings.Split(token[0], "Basic ")[1])
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		} else {
			c.Set("User", ok.ID)
			c.Set("Role", ok.Role)
			c.Next()
		}
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Auth Token Not supplied"})
	return
}

func (h *MiddlewaresImpl) CheckAdminRole(c *gin.Context) {
	if c.GetString("Role") == "admin" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Role not authorized"})
		return
	}
}
