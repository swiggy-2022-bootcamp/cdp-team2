package ports

import (
	"context"

	pb "common/protos/products"
)

type IGRPCHandlers interface {
	GetAvailableProducts(context.Context, *pb.EmptyRequest) (*pb.ProductsResponse, error)
	CheckProductsWithCategory(context.Context, *pb.CategoryIDRequest) (*pb.BoolResponse, error)
	IsProductExists(context.Context, *pb.ProductIDRequest) (*pb.BoolResponse, error)
	GetProductById(context.Context, *pb.ProductIDRequest) (*pb.ProductResponse, error)
	GetProductsByCategoryId(context.Context, *pb.CategoryIDRequest) (*pb.ProductsResponse, error)
	DeductProductsQuantity(context.Context, *pb.CheckoutRequest) (*pb.CheckoutResponse, error)
}
