package ports

import (
	"github.com/gin-gonic/gin"
)

type IMiddlewares interface {
	CheckAuthMiddleware(ctx *gin.Context)
	CheckAdminRole(c *gin.Context)
}
