package ports

import "github.com/gin-gonic/gin"

type IProductsHandlers interface {
	Health(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
	GetProducts(*gin.Context)
	SearchByLimit(*gin.Context)
	SearchByKeyword(*gin.Context)
	SearchByCategoryID(*gin.Context)
	SearchByManufacturerID(*gin.Context)
	SearchByStartPrice(*gin.Context)
}
