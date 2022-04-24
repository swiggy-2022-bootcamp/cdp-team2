package services

import (
	"testing"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
)

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
