package domain

import (
	pb "common/protos/products"
)

type Product struct {
	ID                 int64                 `json:"_id"                 dynamodbav:"_id"`
	Model              string                `json:"model"               dynamodbav:"model"               binding:"required"`
	Quantity           uint                  `json:"quantity"            dynamodbav:"quantity"            binding:"required"`
	Price              string                `json:"price"               dynamodbav:"price"               binding:"required"`
	ManufacturerID     int                   `json:"manufacturer_id"     dynamodbav:"manufacturer_id"     binding:"required"`
	Sku                string                `json:"sku"                 dynamodbav:"sku"                 binding:"required"`
	ProductSeoUrl      *ProductSeoUrl        `json:"product_seo_url"     dynamodbav:"product_seo_url"     binding:"requried"`
	Points             uint                  `json:"points"              dynamodbav:"points"              binding:"requried"`
	Rewards            uint                  `json:"reward"              dynamodbav:"reward"              binding:"requried"`
	Image              string                `json:"image"               dynamodbav:"image"               binding:"requried"`
	ShippingID         int                   `json:"shipping_id"         dynamodbav:"shipping_id"         binding:"requried"`
	Weight             uint                  `json:"weight"              dynamodbav:"weight"              binding:"requried"`
	Length             uint                  `json:"length"              dynamodbav:"length"              binding:"requried"`
	Width              uint                  `json:"width"               dynamodbav:"width"               binding:"requried"`
	Height             uint                  `json:"height"              dynamodbav:"height"              binding:"requried"`
	Minimum            uint                  `json:"minimum"             dynamodbav:"minimum"             binding:"requried"`
	ProductRelated     []int64               `json:"product_related"     dynamodbav:"product_related"     binding:"requried"`
	ProductDescription []*ProductDescription `json:"product_description" dynamodbav:"product_description" binding:"required"`
	ProductCategory    []int64               `json:"product_category"    dynamodbav:"product_category"    binding:"required"`
}

func GetPbProductDescriptions(_productDesc []*ProductDescription) (_pbProductDesc []*pb.ProductDescription) {
	for _, _pDesc := range _productDesc {
		_pbProductDesc = append(_pbProductDesc, _pDesc.GetPbProductDescription())
	}
	return
}

func (p *Product) GetPbProduct() *pb.Product {
	return &pb.Product{
		ID:                 int64(p.ID),
		Model:              p.Model,
		Quantity:           int64(p.Quantity),
		Price:              p.Price,
		ManufacturerID:     int64(p.ManufacturerID),
		Sku:                p.Sku,
		ProductSeoUrl:      p.ProductSeoUrl.GetPbProductSeoUrl(),
		Points:             int64(p.Points),
		Rewards:            int64(p.Rewards),
		Image:              p.Image,
		Weight:             int64(p.Weight),
		Length:             int64(p.Length),
		Width:              int64(p.Width),
		Height:             int64(p.Height),
		Minimum:            int64(p.Minimum),
		ProductRelated:     p.ProductRelated,
		ProductDescription: GetPbProductDescriptions(p.ProductDescription),
		ProductCategory:    p.ProductCategory,
	}
}

func bindProductDesciption(gprcProDes []*pb.ProductDescription) []*ProductDescription {
	var _productDecs []*ProductDescription
	for _, _p := range gprcProDes {
		var _productDes ProductDescription
		_productDecs = append(_productDecs, _productDes.BindGrpcProductDesciption(_p))
	}
	return _productDecs
}

func (p *Product) BindGprcProduct(grpcProduct *pb.Product) {
	p.ID = grpcProduct.ID
	p.Model = grpcProduct.Model
	p.Quantity = uint(grpcProduct.Quantity)
	p.Price = grpcProduct.Price
	p.ManufacturerID = int(grpcProduct.ManufacturerID)
	p.Sku = grpcProduct.Sku
	p.ProductSeoUrl = p.ProductSeoUrl.BindGrpcProductSeoUrl(grpcProduct.ProductSeoUrl)
	p.Points = uint(grpcProduct.Points)
	p.Rewards = uint(grpcProduct.Rewards)
	p.Image = grpcProduct.Image
	p.Weight = uint(grpcProduct.Weight)
	p.Length = uint(grpcProduct.Length)
	p.Width = uint(grpcProduct.Width)
	p.Height = uint(grpcProduct.Height)
	p.Minimum = uint(grpcProduct.Minimum)
	p.ProductRelated = grpcProduct.ProductRelated
	p.ProductDescription = bindProductDesciption(grpcProduct.ProductDescription)
	p.ProductCategory = grpcProduct.ProductCategory
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
