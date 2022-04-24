package services

import (
	"testing"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
)

func TestUpdateCartItemServiceValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetUpdateCartItemService()

	type functionArgs struct {
		product models.Product
	}

	tests := []struct {
		name          string
		args          functionArgs
		expectedError *errors.ServerError
	}{
		{
			name: "ValidateRequestWithNoError",
			args: functionArgs{
				product: models.Product{
					ProductId: "1234",
					Quantity:  5,
				},
			},
			expectedError: nil,
		},
		{
			name: "ValidateRequestWithNoProductId",
			args: functionArgs{
				product: models.Product{
					ProductId: "",
					Quantity:  5,
				},
			},
			expectedError: &errors.ParametersMissingError,
		},
		{
			name: "ValidateRequestWithZeroProductQuantity",
			args: functionArgs{
				product: models.Product{
					ProductId: "123",
					Quantity:  0,
				},
			},
			expectedError: &errors.ParametersMissingError,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := service.ValidateRequest(testCase.args.product)

			if err != testCase.expectedError {
				t.Errorf("Error returned: %v, want: %v", err, testCase.expectedError)
			}
		})
	}
}
