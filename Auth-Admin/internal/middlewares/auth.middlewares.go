package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/auth-admin-service/internal/core/domain"
	usersrv "github.com/auth-admin-service/internal/core/services/usersrv"
	repo "github.com/auth-admin-service/internal/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *MiddlewaresImpl) CheckBasicAuthMiddleware(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) > 0 {
		ok, err := repo.JWTManager.VerifyBasicToken(strings.Split(token[0], "Basic ")[1])
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		} else {
			c.Set("User", ok.ID)
			c.Set("Role", ok.Role)
			c.Next()
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Auth Token Not supplied"})
	}
}

func (h *MiddlewaresImpl) CheckAdminRole(c *gin.Context) {
	if c.GetString("Role") == "admin" {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Role not authorized"})
		return
	}
}

func (h *MiddlewaresImpl) CheckAuthMiddleware(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) > 0 {
		access_token := strings.Split(token[0], "Bearer ")[1]
		ok, err := repo.JWTManager.Verify(access_token)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		} else {
			user := &domain.User{}
			id, err := primitive.ObjectIDFromHex(ok.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error while getting User Id"})
				return
			}
			fmt.Println(id, access_token)
			if err := usersrv.New().FetchUser(&bson.M{"_id": id, "tokens": access_token}, user); err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Token Invalid"})
				return
			} else {
				c.Set("User", ok.ID)
				c.Set("Role", ok.Role)
				c.Next()
			}
		}
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Auth Token Not supplied"})
	}
}
