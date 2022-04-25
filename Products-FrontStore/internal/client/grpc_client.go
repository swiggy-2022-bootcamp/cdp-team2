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
	grpcConn, err := grpc.Dial("http://10.50.2.224:30202", opts...)
	if err != nil {
		log.Fatal(err)
	}
	return pb.NewProductsServicesClient(grpcConn)
}
