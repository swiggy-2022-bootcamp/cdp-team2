package services

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/reward/internal/errors"
)

func TestAddRewardServiceValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetAddRewardService()

	tests := []struct {
		name          string
		expectedError *errors.ServerError
		reward        models.Reward
	}{
		{
			name: "ValidateRequestWithNoError",
			reward: models.Reward{
				CustomerId:  "2324",
				Description: "Reward points Added",
				Points:      21,
			},
			expectedError: nil,
		},
		{
			name: "ValidateRequestWithZeroPoints",
			reward: models.Reward{
				CustomerId:  "2324",
				Description: "Reward points Added",
				Points:      0,
			},
			expectedError: &errors.ParametersMissingError,
		},
		{
			name: "ValidateRequestWithEmptyId",
			reward: models.Reward{
				CustomerId:  "",
				Description: "Reward points Added",
				Points:      23,
			},
			expectedError: &errors.ParametersMissingError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := service.ValidateRequest(testCase.reward)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}

func TestAddRewardServiceProcessRequest(t *testing.T) {
	mockDao = setupTest()
	service := GetAddRewardService()

	tests := []struct {
		name          string
		expectedError *errors.ServerError
		reward        models.Reward
		setupFunc     func()
	}{
		{
			name: "ValidateRequestWithNoError",
			reward: models.Reward{
				CustomerId:  "2324",
				Description: "Reward points Added",
				Points:      21,
			},
			expectedError: nil,
			setupFunc: func() {
				mockDao.On("AddReward", mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "ValidateRequestWithError",
			reward: models.Reward{
				CustomerId:  "2324",
				Description: "Reward points Added",
				Points:      23,
			},
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockDao.On("AddReward", mock.Anything).Return(&errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}

			err := service.ProcessRequest(testCase.reward)
			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
