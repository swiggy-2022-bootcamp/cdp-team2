package grpc

import (
	"context"
	"log"
	"time"

	orderpb "common/protos/order"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	ClientConn
	orderpb.ServiceClient
}

func NewOrderClient(ctx context.Context) (*OrderClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["ORDER_SERVICE"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error creating grpc Order Client : %s", err.Error())
		return nil, err
	}

	log.Printf("[gcon] conn %+v", conn)

	orderClient := orderpb.NewServiceClient(conn)

	ctxx, cancel := context.WithCancel(ctx)

	return &OrderClient{
		ClientConn{c: ctxx, cancel: cancel, DefaultTimeout: time.Hour, gconn: conn},
		orderClient,
	}, nil
}
