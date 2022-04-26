package services

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/literals"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

type AddRewardService interface {
	ValidateRequest(reward models.Reward) *errors.ServerError
	ProcessRequest(reward models.Reward) *errors.ServerError
}

var addRewardServiceStruct AddRewardService
var addRewardServiceOnce sync.Once

type addRewardService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitAddRewardService(config *util.RouterConfig, dao dao.DynamoDAO) AddRewardService {
	addRewardServiceOnce.Do(func() {
		addRewardServiceStruct = &addRewardService{
			config: config,
			dao:    dao,
		}
	})

	return addRewardServiceStruct
}

func GetAddRewardService() AddRewardService {
	if addRewardServiceStruct == nil {
		panic("add reward service not initialized")
	}

	return addRewardServiceStruct
}

func (s *addRewardService) ValidateRequest(reward models.Reward) *errors.ServerError {
	if reward.CustomerId == "" || reward.Description == "" || reward.Points == 0 {
		log.WithFields(log.Fields{
			literals.CustomerId:   reward.CustomerId,
			literals.Description:  reward.Description,
			literals.RewardPoints: reward.Points,
		}).Error(errors.ParametersMissingError.ErrorMessage)
		return &errors.ParametersMissingError
	}

	return nil
}

func (s *addRewardService) ProcessRequest(reward models.Reward) *errors.ServerError {

	err := s.dao.AddReward(reward)
	if err != nil {
		return err
	}

	return nil
}
