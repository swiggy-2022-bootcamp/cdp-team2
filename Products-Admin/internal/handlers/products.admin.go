package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/domain"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/literals"
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
func (h *ProductsHandlers) Health(gctx *gin.Context) {
	contentType := "text/plain; charset=utf-8"
	gctx.Data(http.StatusOK, contentType, []byte(literals.HEALTH_MESSAGE))
}

func (h *ProductsHandlers) GetProducts(gctx *gin.Context) {
	_products, err := h.ProductsServices.GetProducts()
	if err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": _products})
}

func (h *ProductsHandlers) AddProduct(gctx *gin.Context) {
	var _product domain.Product
	if err := gctx.ShouldBindJSON(&_product); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	productID, err := h.ProductsServices.AddProduct(&_product)
	if err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprint(productID) + " Product Added."})
}

func (h *ProductsHandlers) UpdateProduct(gctx *gin.Context) {
	productIDStr := gctx.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid product id."})
		return
	}
	var _product domain.Product
	if err := gctx.ShouldBindJSON(&_product); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid product information."})
		return
	}
	if err := h.ProductsServices.UpdateProduct(int64(productID), &_product); err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": "Product Updated."})
}

func (h *ProductsHandlers) DeleteProduct(gctx *gin.Context) {
	productIDStr := gctx.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid product id."})
		return
	}
	if err := h.ProductsServices.DeleteProduct(int64(productID)); err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusAccepted, gin.H{"message": "Product deleted."})
}
