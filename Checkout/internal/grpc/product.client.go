package grpc

import (
	"context"
	"log"
	"time"

	productpb "common/protos/products"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
	"google.golang.org/grpc"
)

type ProductClient struct {
	ClientConn
	productpb.ProductsServicesClient
}

func NewProductClient(ctx context.Context) (*ProductClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["PRODUCT_SERVICE"], grpc.WithInsecure())
	if err != nil {
		log.Printf("Error creating grpc product Client : %s", err.Error())
		return nil, err
	}

	productClient := productpb.NewProductsServicesClient(conn)

	ctxx, cancel := context.WithCancel(ctx)

	return &ProductClient{
		ClientConn{c: ctxx, cancel: cancel, DefaultTimeout: time.Hour, gconn: conn},
		productClient,
	}, nil
}
