package ports

import "github.com/gin-gonic/gin"

type AuthService interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
}
