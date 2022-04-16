package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type DeleteCartItemService interface {
	ValidateRequest() *errors.ServerError
	ProcessRequest() *errors.ServerError
}

var deleteCartItemServiceStruct DeleteCartItemService
var deleteCartItemServiceOnce sync.Once

type deleteCartItemService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitDeleteCartItemService(config *util.RouterConfig, dao dao.DynamoDAO) DeleteCartItemService {
	deleteCartItemServiceOnce.Do(func() {
		deleteCartItemServiceStruct = &deleteCartItemService{
			config: config,
			dao:    dao,
		}
	})

	return deleteCartItemServiceStruct
}

func GetDeleteCartItemService() DeleteCartItemService {
	if deleteCartItemServiceStruct == nil {
		panic("delete cart item service is not initialized")
	}

	return deleteCartItemServiceStruct
}

func (s *deleteCartItemService) ValidateRequest() *errors.ServerError {
	return nil
}

func (s *deleteCartItemService) ProcessRequest() *errors.ServerError {
	return nil
}
