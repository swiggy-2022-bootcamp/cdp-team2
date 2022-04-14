package ports

import (
	"github.com/gin-gonic/gin"
)

type IHandlers interface {
	Health(*gin.Context)
	Login(*gin.Context)
	OAuth(*gin.Context)
}
