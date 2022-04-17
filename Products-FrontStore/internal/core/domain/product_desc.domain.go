package domain

import (
	pb "common/pkg/protos"
)

type ProductDescription struct {
	LanguageID      int64  `json:"language_id"      bson:"language_id"`
	Name            string `json:"name"             bson:"name"`
	Description     string `json:"description"      bson:"description"`
	MetaTitle       string `json:"meta_title"       bson:"meta_title"`
	MetaDescription string `json:"meta_description" bson:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"     bson:"meta_keyword"`
	Tag             string `json:"tag"              bson:"tag"`
}

func (p *ProductDescription) GetPbProductDescription() *pb.ProductDescription {
	return &pb.ProductDescription{
		LanguageID:     int64(p.LanguageID),
		Name:           p.Name,
		Description:    p.Description,
		MetaTitle:      p.MetaTitle,
		MetaDescriptio: p.MetaDescription,
		MetaKeyword:    p.MetaKeyword,
		Tag:            p.Tag,
	}
}

func (p *ProductDescription) BindGrpcProductDesciption(grpcProductDesc *pb.ProductDescription) *ProductDescription {
	return &ProductDescription{
		LanguageID:      grpcProductDesc.LanguageID,
		Name:            grpcProductDesc.Name,
		Description:     grpcProductDesc.Description,
		MetaTitle:       grpcProductDesc.MetaTitle,
		MetaDescription: grpcProductDesc.Description,
		MetaKeyword:     grpcProductDesc.MetaKeyword,
		Tag:             grpcProductDesc.Tag,
	}
}
