package domain

import (
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/protos"
)

type ProductSeoUrl struct {
	Keyword    string `json:"keyword"     dynamodbav:"keyword"`
	LanguageID int64  `json:"language_id" dynamodbav:"langauge_id"`
	StoreID    int64  `json:"store_id"    dynamodbav:"store_id"`
}

func (p *ProductSeoUrl) GetPbProductSeoUrl() *pb.ProductSeoUrl {
	return &pb.ProductSeoUrl{
		Keyword:    p.Keyword,
		LanguageID: int64(p.LanguageID),
		StoreID:    int64(p.StoreID),
	}
}

func (p *ProductSeoUrl) BindGrpcProductSeoUrl(grpcProductsSeoUrl *pb.ProductSeoUrl) *ProductSeoUrl {
	return &ProductSeoUrl{
		Keyword:    grpcProductsSeoUrl.Keyword,
		LanguageID: grpcProductsSeoUrl.LanguageID,
		StoreID:    grpcProductsSeoUrl.StoreID,
	}
}
