package domain

import (
	pb "common/protos/products"
)

type ProductDescription struct {
	LanguageID      int64  `json:"language_id"      dynamodbav:"language_id"`
	Name            string `json:"name"             dynamodbav:"name"`
	Description     string `json:"description"      dynamodbav:"description"`
	MetaTitle       string `json:"meta_title"       dynamodbav:"meta_title"`
	MetaDescription string `json:"meta_description" dynamodbav:"meta_description"`
	MetaKeyword     string `json:"meta_keyword"     dynamodbav:"meta_keyword"`
	Tag             string `json:"tag"              dynamodbav:"tag"`
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
