package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

// GetUser godoc
// @Summary      Get All Products
// @Description  products get API
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200  {object} 	responses.ProductsMessage
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       / [get]
func (h *ProductsHandlers) GetProducts(gctx *gin.Context) {
	_products, err := h.ProductsServices.GetProducts()
	if err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": _products})
}

// GetUser godoc
// @Summary      Add Products
// @Description  product post API
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product body requests.ProductRequest  true "product request structure"
// @Success      200  {object} 	responses.MessageResponse
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       / [post]
func (h *ProductsHandlers) AddProduct(gctx *gin.Context) {
	var _product domain.Product
	if err := gctx.ShouldBindJSON(&_product); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	productID, err2 := h.ProductsServices.AddProduct(&_product)
	if err2 != nil {
		gctx.JSON(err2.GetErrCode(), gin.H{"message": err2.Error()})
		return
	}
	gctx.JSON(http.StatusCreated, gin.H{"message": fmt.Sprint(productID) + " Product Added."})
}

// GetUser godoc
// @Summary      Update Products
// @Description  product put API
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id path string true "produtID"
// @Param        product body requests.ProductIDRequest  true "product request structure"
// @Success      200  {object} 	responses.MessageResponse
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /{id} [put]
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

// GetUser godoc
// @Summary      Delete Products
// @Description  product delete API
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id path string true "produtID"
// @Success      200  {object} 	responses.MessageResponse
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /{id} [delete]
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
	gctx.JSON(http.StatusOK, gin.H{"message": "Product deleted."})
}

func (ph *ProductsHandlers) SearchByLimit(gctx *gin.Context) {
	gctx.JSON(http.StatusOK, nil)
}

// GetUser godoc
// @Summary      Search Products by Keyword
// @Description  product search API
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        keyword path string true "search keyword"
// @Success      200  {object} 	responses.ProductsMessage
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /search/keyword/{keyword} [get]
func (ph *ProductsHandlers) SearchByKeyword(gctx *gin.Context) {

	// Extract the keyword from the path params
	keyword := strings.ToLower(gctx.Param("keyword"))

	/*
	* 	Get all the products that contains the keyword
	*	in their product name, model name, sku
	 */
	products, err := ph.ProductsServices.SearchByKeyword(keyword)
	if err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": products})
}

// GetUser godoc
// @Summary      Search Products by CategoryID
// @Description  product search API
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        id path string true "categoryID"
// @Success      200  {object} 	responses.ProductsMessage
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /search/category/{id} [get]
func (ph *ProductsHandlers) SearchByCategoryID(gctx *gin.Context) {

	// Extract the category ID
	categoryIDStr := gctx.Param("id")

	// Convert the string type to int type
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid category id."})
		return
	}

	products, err2 := ph.ProductsServices.GetProductsByCategoryId(int64(categoryID))
	if err2 != nil {
		gctx.JSON(err2.GetErrCode(), gin.H{"message": err2.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": products})
}

// GetUser godoc
// @Summary      Search Products by ManufacturerID
// @Description  product search API
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        id path string true "manufacturerID"
// @Success      200  {object} 	responses.ProductsMessage
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /search/manufacturer/{id} [get]
func (ph *ProductsHandlers) SearchByManufacturerID(gctx *gin.Context) {

	// Extract the manufacturer ID
	manufacturerIDStr := gctx.Param("id")

	// Convert the string type to int type
	manufacturerID, err := strconv.Atoi(manufacturerIDStr)
	if err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide valid manufacturer id."})
		return
	}

	products, err2 := ph.ProductsServices.SearchByManufacturerID(int64(manufacturerID))
	if err2 != nil {
		gctx.JSON(err2.GetErrCode(), gin.H{"message": err2.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": products})
}

// GetUser godoc
// @Summary      Search Products by Starting Price
// @Description  product search API
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param        price path string true "Starting Price"
// @Success      200  {object} 	responses.ProductsMessage
// @Failure      400  {object} 	responses.MessageResponse
// @Failure      500  {object} 	responses.MessageResponse
// @Failure      502  {object} 	responses.MessageResponse
// @Router       /search/start/{price} [get]
func (ph *ProductsHandlers) SearchByStartPrice(gctx *gin.Context) {
	price := gctx.Param("price")

	products, err := ph.ProductsServices.SearchByStartPrice(price)
	if err != nil {
		gctx.JSON(err.GetErrCode(), gin.H{"message": err.Error()})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{"message": products})
}
