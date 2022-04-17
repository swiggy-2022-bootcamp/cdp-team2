package client

import (
	"log"

	pb "common/pkg/protos"

	"google.golang.org/grpc"
)

type ProductsGRPCClient pb.ProductsServicesClient

func NewProductsGRPCClient() ProductsGRPCClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	grpcConn, err := grpc.Dial("localhost:8002", opts...)
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewProductsServicesClient(grpcConn)
}
