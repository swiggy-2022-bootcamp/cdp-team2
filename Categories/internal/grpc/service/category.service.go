package service

import (
	pb "common/protos/category"
	"context"
	"fmt"
	"log"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/services"
)

type CategoryGrpcServer struct {
	pb.UnimplementedCategoryServiceServer
	service services.IService
}

func NewCategoryGrpcServer() (*CategoryGrpcServer, error) {
	s, err := services.NewCategoryService()
	return &CategoryGrpcServer{
		service: s,
	}, err
}

func (cgs *CategoryGrpcServer) GetCategory(ctx context.Context, in *pb.GetCategoryInput) (*pb.Category, error) {
	log.Printf("[Info] grpc call for get category %d", in.CategoryId)

	if in.CategoryId == 0 {
		return nil, fmt.Errorf("category id cannot be 0")
	}

	cat, err := cgs.service.GetByID(int(in.CategoryId))

	if err != nil {
		return nil, err
	}

	return &pb.Category{
		CategoryId: int32(cat.CategoryId),
		CategoryDescription: &pb.CategoryDesc{
			Name:            cat.CategoryDesc.Name,
			Description:     cat.CategoryDesc.Description,
			MetaDescription: cat.CategoryDesc.MetaDescription,
			MetaKeyword:     cat.CategoryDesc.MetaKeyword,
			MetaTitle:       cat.CategoryDesc.MetaTitle,
		},
	}, nil
}
