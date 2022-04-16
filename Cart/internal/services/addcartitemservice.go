package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type AddCartItemService interface {
	ValidateRequest() *errors.ServerError
	ProcessRequest(product models.Product) *errors.ServerError
}

var addCartItemServiceStruct AddCartItemService
var addCartItemServiceOnce sync.Once

type addCartItemService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitAddCartItemService(config *util.RouterConfig, dao dao.DynamoDAO) AddCartItemService {
	addCartItemServiceOnce.Do(func() {
		addCartItemServiceStruct = &addCartItemService{
			config: config,
			dao:    dao,
		}
	})

	return addCartItemServiceStruct
}

func GetAddCartItemService() AddCartItemService {
	if addCartItemServiceStruct == nil {
		panic("add cart item service is not initialized")
	}

	return addCartItemServiceStruct
}

func (s *addCartItemService) ValidateRequest() *errors.ServerError {
	return nil
}

func (s *addCartItemService) ProcessRequest(product models.Product) *errors.ServerError {
	return s.dao.AddCartItem(product)
}
