package requests

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-FrontStore/internal/core/domain"
)

type ProductIDRequest struct {
	Product *domain.Product `json:"product"`
}
type ProductRequest struct {
	Model              string                       `json:"model"               dynamodbav:"model"               binding:"required"`
	Quantity           uint                         `json:"quantity"            dynamodbav:"quantity"            binding:"required"`
	Price              string                       `json:"price"               dynamodbav:"price"               binding:"required"`
	ManufacturerID     int                          `json:"manufacturer_id"     dynamodbav:"manufacturer_id"     binding:"required"`
	Sku                string                       `json:"sku"                 dynamodbav:"sku"                 binding:"required"`
	ProductSeoUrl      *domain.ProductSeoUrl        `json:"product_seo_url"     dynamodbav:"product_seo_url"     binding:"required"`
	Points             uint                         `json:"points"              dynamodbav:"points"              binding:"required"`
	Rewards            uint                         `json:"reward"              dynamodbav:"reward"              binding:"required"`
	Image              string                       `json:"image"               dynamodbav:"image"               binding:"required"`
	ShippingID         int                          `json:"shipping_id"         dynamodbav:"shipping_id"         binding:"required"`
	Weight             uint                         `json:"weight"              dynamodbav:"weight"              binding:"required"`
	Length             uint                         `json:"length"              dynamodbav:"length"              binding:"required"`
	Width              uint                         `json:"width"               dynamodbav:"width"               binding:"required"`
	Height             uint                         `json:"height"              dynamodbav:"height"              binding:"required"`
	Minimum            uint                         `json:"minimum"             dynamodbav:"minimum"             binding:"required"`
	ProductRelated     []int64                      `json:"product_related"     dynamodbav:"product_related"     binding:"required"`
	ProductDescription []*domain.ProductDescription `json:"product_description" dynamodbav:"product_description" binding:"required"`
	ProductCategory    []int64                      `json:"product_category"    dynamodbav:"product_category"    binding:"required"`
}
