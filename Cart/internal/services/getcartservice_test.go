package services

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
)

func TestGetCartServiceProcessRequest(t *testing.T) {
	mockDao := setupTest()
	service := GetGetCartService()

	tests := []struct {
		name          string
		customerId    string
		expectedError *errors.ServerError
		setupFunc     func()
	}{
		{
			name:          "ProcessRequestWithNoError",
			customerId:    "123",
			expectedError: nil,
			setupFunc: func() {
				mockDao.On("GetCart", mock.Anything).Return(models.Cart{}, nil).Once()
			},
		},
		{
			name:          "ProcessRequestWithError",
			customerId:    "123",
			expectedError: &errors.InternalError,
			setupFunc: func() {
				mockDao.On("GetCart", mock.Anything).Return(models.Cart{}, &errors.InternalError).Once()
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setupFunc != nil {
				testCase.setupFunc()
			}

			_, err := service.ProcessRequest(testCase.customerId)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
