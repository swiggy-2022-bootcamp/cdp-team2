package grpchandlers

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/reward/protos"
)

func (s *Server) RedeemReward(ctx context.Context, reward *pb.CalculateRewardRequest) (*pb.CalculateRewardResponse, error) {

	// call to Order MS to get list of products

	// calculation of total reward points that customer wants to use

	// call to Order MS to update the finalPrice of the order

	// return the final response to calling microservice
	return nil, nil
}

func (s *Server) CalculateReward(ctx context.Context, reward *pb.CalculateRewardRequest) (*pb.CalculateRewardResponse, error) {
	return nil, nil
}
