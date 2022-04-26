package grpc

import (
	"context"
	"log"
	"time"

	authpb "common/protos/authpb"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
	ClientConn
	authpb.AuthServiceClient
}

func NewAuthClient(ctx context.Context) (*AuthClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["AUTH_SERVICE"], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error creating grpc auth Client : %s", err.Error())
		return nil, err
	}

	auth := authpb.NewAuthServiceClient(conn)

	ctxx, cancel := context.WithCancel(ctx)

	return &AuthClient{
		ClientConn{c: ctxx, cancel: cancel, DefaultTimeout: time.Hour, gconn: conn},
		auth,
	}, nil
}
