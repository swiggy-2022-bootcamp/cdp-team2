package ports

import (
	pb "common/protos/products"
)

type IProductsGrpcClientServices interface {
	pb.ProductsServicesClient
}
