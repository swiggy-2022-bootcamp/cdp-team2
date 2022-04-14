package ports

import "github.com/gin-gonic/gin"

type IProdcutsHandlers interface {
	Health(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	GetProducts(*gin.Context)
}
