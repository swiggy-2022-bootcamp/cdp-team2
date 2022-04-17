package domain

import (
	pb "common/pkg/protos"
)

type ProductSeoUrl struct {
	Keyword    string `json:"keyword" bson:"keyword"`
	LanguageID int64  `json:"language_id" bson:"langauge_id"`
	StoreID    int64  `json:"store_id" bson:"store_id"`
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
