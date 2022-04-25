package client

import (
	"context"
	"log"
	"time"

	productpb "common/protos/products"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClientConn interface {
	// Start() error
	CtxWithTimeOut() context.Context
	Stop() error
}

type ClientConn struct {
	c              context.Context
	cancel         context.CancelFunc
	DefaultTimeout time.Duration
	gconn          *grpc.ClientConn
}

func (cc *ClientConn) CtxWithTimeOut() context.Context {
	ctx, _ := context.WithTimeout(cc.c, cc.DefaultTimeout)
	return ctx
}

func (cc *ClientConn) Stop() {
	cc.cancel()
	cc.gconn.Close()
	log.Println("Stopping grpc client")
}

type ProductClient struct {
	ClientConn
	productpb.ProductsServicesClient
}

func NewProductClient(ctx context.Context) (*ProductClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["PRODUCT_SERVICE"], grpc.WithTransportCredentials(insecure.NewCredentials()))
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
