package grpc_handlers

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/protos"
)

type GRPCHandlers struct {
	pb.UnimplementedProductsServicesServer
	ProductsServices ports.IProductsServices
}

func NewGRPCServer(productsServices ports.IProductsServices) *GRPCHandlers {
	return &GRPCHandlers{
		ProductsServices: productsServices,
	}
}

var _ ports.IGRPCHandlers = (*GRPCHandlers)(nil)
var _ pb.ProductsServicesServer = (*GRPCHandlers)(nil)

func (gh *GRPCHandlers) GetAvailableProducts(ctx context.Context, emptyReq *pb.EmptyRequest) (*pb.ProductsResponse, error) {
	_products, err := gh.ProductsServices.GetProducts()
	var _pbProducts []*pb.Product
	if err != nil {
		return nil, err
	}
	for _, _p := range _products {
		_pbProducts = append(_pbProducts, _p.GetPbProduct())
	}
	return &pb.ProductsResponse{
		Products: _pbProducts,
	}, nil
}

func (gh *GRPCHandlers) CheckProductsWithCategory(ctx context.Context, categoryID *pb.CategoryIDRequest) (*pb.BoolResponse, error) {
	return nil, nil
}
