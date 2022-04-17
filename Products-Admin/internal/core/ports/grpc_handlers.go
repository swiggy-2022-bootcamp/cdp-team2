package ports

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/protos"
)

type IGRPCHandlers interface {
	GetAvailableProducts(context.Context, *pb.EmptyRequest) (*pb.ProductsResponse, error)
	CheckProductsWithCategory(context.Context, *pb.CategoryIDRequest) (*pb.BoolResponse, error)
}
