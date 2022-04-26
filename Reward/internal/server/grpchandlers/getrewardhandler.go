package grpchandlers

import (
	"context"
	"errors"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/reward/protos"

	log "github.com/sirupsen/logrus"
)

type Server struct {
}

func (s *Server) GetReward(ctx context.Context, reward *pb.RewardRequest) (*pb.RewardResponse, error) {
	dao := dao.GetDynamoDAO()

	customerId := reward.CustomerId
	rewardModel, err := dao.GetReward(customerId)
	if err != nil {
		log.Error("Error while getting the reward: ", err.ErrorMessage)
		return nil, errors.New("error occurred")
	}

	return &pb.RewardResponse{
		Reward: &pb.Reward{
			CustomerId:  customerId,
			Points:      rewardModel.Points,
			Description: rewardModel.Description,
		},
	}, nil
}
