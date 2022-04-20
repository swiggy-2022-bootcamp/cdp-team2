package grpc_handlers

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/protos"
)

type GRPCHandlers struct {
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

func (gh *GRPCHandlers) GetProductById(ctx context.Context, productIDReq *pb.ProductIDRequest) (*pb.ProductResponse, error) {
	_product, err := gh.ProductsServices.GetProductById(productIDReq.ProductID)
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{
		Product: _product.GetPbProduct(),
	}, nil
}

func (gh *GRPCHandlers) CheckProductsWithCategory(ctx context.Context, categoryID *pb.CategoryIDRequest) (*pb.BoolResponse, error) {
	return nil, nil
}

func (gh *GRPCHandlers) IsProductExists(ctx context.Context, productIDReq *pb.ProductIDRequest) (*pb.BoolResponse, error) {
	exists, err := gh.ProductsServices.IsProductExists(productIDReq.ProductID)
	if err != nil {
		return &pb.BoolResponse{
			Exists: false,
		}, err
	}
	if exists == false {
		return &pb.BoolResponse{
			Exists: false,
		}, nil
	}
	return &pb.BoolResponse{
		Exists: true,
	}, nil
}

func (gh *GRPCHandlers) GetProductsByCategoryId(ctx context.Context, categoryIDReq *pb.CategoryIDRequest) (*pb.ProductsResponse, error) {
	_products, err := gh.ProductsServices.GetProductsByCategoryId(categoryIDReq.CategoryID)
	if err != nil {
		return nil, err
	}

	// Convert domain Products to proto Products
	var _pbProducts []*pb.Product
	for _, _product := range _products {
		_pbProducts = append(_pbProducts, _product.GetPbProduct())
	}

	return &pb.ProductsResponse{
		Products: _pbProducts,
	}, nil
}
