package grpc_handlers

import (
	"context"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/internal/core/ports"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Products-Admin/pkg/protos"
)

type GRPCHandlers struct {
}

func NewGRPCServer(ctx context.Context) *GRPCHandlers {
	return &GRPCHandlers{}
}

var _ ports.IGRPCHandlers = (*GRPCHandlers)(nil)

func (gh *GRPCHandlers) GetAvailableProducts(ctx context.Context, emptyReq *pb.EmptyRequest) (*pb.ProductsResponse, error) {
	return nil, nil
}

func (gh *GRPCHandlers) CheckProductsWithCategory(ctx context.Context, categoryID *pb.CategoryIDRequest) (*pb.BoolResponse, error) {
	return nil, nil
}
