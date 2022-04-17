package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/literals"
)

type ProductsHandlers struct {
	ProductsServices ports.IProductsServices
}

var _ ports.IProductsHandlers = (*ProductsHandlers)(nil)

func NewHandlers(productsServices ports.IProductsServices) *ProductsHandlers {
	return &ProductsHandlers{
		ProductsServices: productsServices,
	}
}

// Health godoc
// @Summary      Health Check Route
// @Description  API to check products-admin health
// @Tags         Health
// @Accept       json
// @Produce      plain/text
// @Success      200  {object} 	string
// @Router       / [get]
func (ph *ProductsHandlers) Health(gctx *gin.Context) {
	contentType := "text/plain; charset=utf-8"
	gctx.Data(http.StatusOK, contentType, []byte(literals.HEALTH_MESSAGE))
}

func (ph *ProductsHandlers) GetProductList(gctx *gin.Context) {
	_products, err := ph.ProductsServices.GetProductList()
	if err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": _products})
}

func (ph *ProductsHandlers) GetProductById(gctx *gin.Context) {
	productIDStr := gctx.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid product id."})
		return
	}
	product, Err := ph.ProductsServices.GetProductById(int64(productID))
	if Err != nil {
		gctx.JSON(Err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"product": product})
}
