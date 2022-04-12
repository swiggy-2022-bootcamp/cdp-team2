package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

type HealthCheckService interface {
	ProcessRequest() models.HealthCheck
}

var healthCheckServiceStruct HealthCheckService
var healthCheckServiceOnce sync.Once

type healthCheckService struct {
	config *util.RouterConfig
}

func InithHealthCheckService(config *util.RouterConfig) HealthCheckService {
	healthCheckServiceOnce.Do(func() {
		healthCheckServiceStruct = &healthCheckService{
			config: config,
		}
	})

	return healthCheckServiceStruct
}

func GetHealthCheckService() HealthCheckService {
	if healthCheckServiceStruct == nil {
		panic("health check service not initialised")
	}

	return healthCheckServiceStruct
}

func (s *healthCheckService) ProcessRequest() models.HealthCheck {
	return models.HealthCheck{Status: "UP"}
}
