package grpc

import (
	"context"
	"log"

	orderpb "common/protos/order"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
	"google.golang.org/grpc"
)

type IClientConn interface {
	// Start() error
	Stop() error
}

type ClientConn struct {
	c      context.Context
	cancel context.CancelFunc
	// DefaultTimeout time.Duration
	gconn *grpc.ClientConn
}

func (cc *ClientConn) Stop() {
	cc.cancel()
	cc.gconn.Close()
	log.Println("Stopping grpc client")
}

type OrderClient struct {
	ClientConn
	orderpb.OrderServiceClient
}

func NewOrderClient(ctx context.Context) (*OrderClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["ORDER_SERVICE"], grpc.WithInsecure())
	if err != nil {
		log.Printf("Error creating grpc Order Client : %s", err.Error())
		return nil, err
	}

	orderClient := orderpb.NewOrderServiceClient(conn)

	ctxx, cancel := context.WithCancel(ctx)

	return &OrderClient{
		ClientConn{ctxx, cancel, conn},
		orderClient,
	}, nil
}
