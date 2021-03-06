package client

import (
	"log"

	pb "common/protos/products"

	"google.golang.org/grpc"
)

type ProductsGRPCClient pb.ProductsServicesClient

func NewProductsGRPCClient() ProductsGRPCClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	grpcConn, err := grpc.Dial("35.84.28.237:30202", opts...)
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewProductsServicesClient(grpcConn)
}
