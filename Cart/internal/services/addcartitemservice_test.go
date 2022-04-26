package services

import (
	"sync"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/config"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/mocks"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
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
		InitAddCartItemService(routerConfig, mockDao)
		InitDeleteCartItemService(routerConfig, mockDao)
		InitUpdateCartItemService(routerConfig, mockDao)
		InitEmptyCartService(routerConfig, mockDao)
		InitHealthCheckService(routerConfig)
		InitGetCartService(routerConfig, mockDao)

	})

	return mockDao
}

func TestAddCartItemServiceValidateRequest(t *testing.T) {
	_ = setupTest()
	service := GetAddCartItemService()

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
