package ports

import "github.com/gin-gonic/gin"

type IProductsHandlers interface {
	Health(*gin.Context)
	GetProductList(*gin.Context)
	GetProductById(*gin.Context)
}
