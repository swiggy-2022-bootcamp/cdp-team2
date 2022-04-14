package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

type AddRewardService interface {
	ProcessRequest()
}

var addRewardServiceStruct AddRewardService
var addRewardServiceOnce sync.Once

type addRewardService struct {
	config *util.RouterConfig
}

func InitAddRewardService(config *util.RouterConfig) AddRewardService {
	addRewardServiceOnce.Do(func() {
		addRewardServiceStruct = &addRewardService{
			config: config,
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

func (s *addRewardService) ProcessRequest() {
}
