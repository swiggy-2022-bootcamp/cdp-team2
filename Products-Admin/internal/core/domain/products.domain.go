package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID                 primitive.ObjectID    `json:"_id"                 bson:"_id"`
	Model              string                `json:"model"               bson:"model"`
	Quantity           uint                  `json:"quantity"            bson:"quantity"`
	Price              string                `json:"price"               bson:"price"`
	ManufacturerID     primitive.ObjectID    `json:"manufacturer_id"     bson:"manufacturer_id"`
	Sku                string                `json:"sku"                 bson:"sku"`
	ProductSeoUrl      []*ProductSeoUrl      `json:"product_seo_url"     bson:"product_seo_url"`
	Points             uint                  `json:"points"              bson:"points"`
	Rewards            uint                  `json:"reward"              bson:"reward"`
	Image              string                `json:"image"               bson:"image"`
	ShippingID         primitive.ObjectID    `json:"shipping_id"         bson:"shipping_id"`
	Weight             uint                  `json:"weight"              bson:"weight"`
	Length             uint                  `json:"length"              bson:"length"`
	Width              uint                  `json:"width"               bson:"width"`
	Height             uint                  `json:"height"              bson:"height"`
	Minimum            uint                  `json:"minimum"             bson:"minimum"`
	ProductRelated     []primitive.ObjectID  `json:"product_related"     bson:"product_related"`
	ProductDescription []*ProductDescription `json:"product_description" bson:"product_description"`
	ProductCategory    []primitive.ObjectID  `json:"product_category"    bson:"product_category"`
}

func (p *Product) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":                 p.ID,
		"model":               p.Model,
		"quantity":            p.Quantity,
		"manafacturer_id":     p.ManufacturerID,
		"sku":                 p.Sku,
		"product_seo_url":     p.ProductSeoUrl,
		"points":              p.Points,
		"rewards":             p.Rewards,
		"image":               p.Image,
		"shipping_id":         p.ShippingID,
		"weight":              p.Weight,
		"length":              p.Length,
		"width":               p.Width,
		"height":              p.Height,
		"minimum":             p.Minimum,
		"product_related":     p.ProductRelated,
		"product_category":    p.ProductCategory,
		"product_description": p.ProductDescription,
	}
}
