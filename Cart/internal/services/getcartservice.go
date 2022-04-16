package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type GetCartService interface {
	ProcessRequest() (models.Cart, *errors.ServerError)
}

var getCartServiceStruct GetCartService
var getCartServiceOnce sync.Once

type getCartService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitGetCartService(config *util.RouterConfig, dao dao.DynamoDAO) GetCartService {
	getCartServiceOnce.Do(func() {
		getCartServiceStruct = &getCartService{
			config: config,
			dao:    dao,
		}
	})

	return getCartServiceStruct
}

func GetGetCartService() GetCartService {
	if getCartServiceStruct == nil {
		panic("get cart service is not initialized")
	}

	return getCartServiceStruct
}

func (s *getCartService) ProcessRequest() (models.Cart, *errors.ServerError) {
	return s.dao.GetCart("12233")
}
