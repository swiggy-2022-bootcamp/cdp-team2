package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type EmptyCartService interface {
	ProcessRequest(customerId string) *errors.ServerError
}

var emptyCartServiceStruct EmptyCartService
var emptyCartServiceOnce sync.Once

type emptyCartService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitEmptyCartService(config *util.RouterConfig, dao dao.DynamoDAO) EmptyCartService {
	emptyCartServiceOnce.Do(func() {
		emptyCartServiceStruct = &emptyCartService{
			config: config,
			dao:    dao,
		}
	})

	return emptyCartServiceStruct
}

func GetEmptyCartService() EmptyCartService {
	if emptyCartServiceStruct == nil {
		panic("empty cart service is not initialized")
	}

	return emptyCartServiceStruct
}

func (s *emptyCartService) ProcessRequest(customerId string) *errors.ServerError {
	return s.dao.EmptyCart(customerId)
}
