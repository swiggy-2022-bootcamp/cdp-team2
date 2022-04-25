package services

import (
	"sync"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/util"
)

var testSync sync.Once
var mockDao *mocks.DynamoDAO

func setupTest() *mocks.DynamoDAO {
	webserverconfig, err := config.FromEnv()
	if err != nil {
		log.WithError(err).Error("error occurred while reading the web configurations")
	}

	routerConfig := &util.RouterConfig{
		WebServerConfig: webserverconfig,
	}

	testSync.Do(func() {
		mockDao = &mocks.DynamoDAO{}
		InitHealthCheckService(routerConfig)
		InitAddRewardService(routerConfig, mockDao)
	})

	return mockDao
}

func TestHealthCheckServiceProcessRequest(t *testing.T) {
	_ = setupTest()
	service := GetHealthCheckService()

	tests := []struct {
		name           string
		expectedStatus models.HealthCheck
	}{
		{
			name:           "ProcessRequestWithNoError",
			expectedStatus: models.HealthCheck{Status: "UP"},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			status := service.ProcessRequest()

			if status != testCase.expectedStatus {
				t.Errorf("Status returned: %v, want: %v", status, testCase.expectedStatus)
			}
		})
	}
}
