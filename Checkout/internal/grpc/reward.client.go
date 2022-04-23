package grpc

import (
	"context"
	"log"
	"time"

	rewardpb "common/protos/reward"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Checkout/config"
	"google.golang.org/grpc"
)

type RewardClient struct {
	ClientConn
	rewardpb.RewardServiceClient
}

func NewRewardClient(ctx context.Context) (*RewardClient, error) {
	conn, err := grpc.Dial(config.GrpcAdd["REWARD_SERVICE"], grpc.WithInsecure())
	if err != nil {
		log.Printf("Error creating grpc product Client : %s", err.Error())
		return nil, err
	}

	rewardClient := rewardpb.NewRewardServiceClient(conn)

	ctxx, cancel := context.WithCancel(ctx)

	return &RewardClient{
		ClientConn{c: ctxx, cancel: cancel, DefaultTimeout: time.Hour, gconn: conn},
		rewardClient,
	}, nil
}
