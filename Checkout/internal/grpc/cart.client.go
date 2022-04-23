package grpc

import (
	"context"
	"log"
	"time"

	cartpb "common/protos/cart"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
	"google.golang.org/grpc"
)

type CartClient struct {
	ClientConn
	cartpb.CartServiceClient
}

func NewCartClient(ctx context.Context) (*CartClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["CART_SERVICE"], grpc.WithInsecure())
	if err != nil {
		log.Printf("Error creating grpc cart Client : %s", err.Error())
		return nil, err
	}

	cartClient := cartpb.NewCartServiceClient(conn)

	ctxx, cancel := context.WithCancel(ctx)

	return &CartClient{
		ClientConn{c: ctxx, cancel: cancel, DefaultTimeout: time.Hour, gconn: conn},
		cartClient,
	}, nil
}
