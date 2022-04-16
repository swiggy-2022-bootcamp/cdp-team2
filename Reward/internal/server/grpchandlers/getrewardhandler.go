package grpchandlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/reward/protos"
	"google.golang.org/protobuf/proto"

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

	resBodyBytes := new(bytes.Buffer)
	goErr := json.NewEncoder(resBodyBytes).Encode(rewardModel)
	if goErr != nil {
		log.WithError(goErr).Error("error occurred while encoding the reward model")
		return nil, goErr
	}

	rewardpb := &pb.RewardResponse{}
	proto.Unmarshal(resBodyBytes.Bytes(), rewardpb)

	return rewardpb, nil
}
