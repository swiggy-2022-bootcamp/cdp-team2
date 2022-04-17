package grpchandlers

import (
	"context"

	pb "github.com/swiggy-2022-bootcamp/cdp-team2/reward/protos"
)

func (s *Server) RedeemReward(ctx context.Context, reward *pb.CalculateRewardRequest) (*pb.CalculateRewardResponse, error) {

	return nil, nil
}

func (s *Server) CalculateReward(ctx context.Context, reward *pb.CalculateRewardRequest) (*pb.CalculateRewardResponse, error) {
	return nil, nil
}
