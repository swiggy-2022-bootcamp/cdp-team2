package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type UpdateCartItemService interface {
	ValidateRequest() *errors.ServerError
	ProcessRequest() *errors.ServerError
}

var updateCartItemServiceStruct UpdateCartItemService
var updateCartItemServiceOnce sync.Once

type updateCartService struct {
	config *util.RouterConfig
	//dao    mongodao.MongoDAO
}

func InitUpdateCartItemService(config *util.RouterConfig) UpdateCartItemService {
	updateCartItemServiceOnce.Do(func() {
		updateCartItemServiceStruct = &updateCartService{
			config: config,
		}
	})

	return updateCartItemServiceStruct
}

func GetUpdateCartItemService() UpdateCartItemService {
	if updateCartItemServiceStruct == nil {
		panic("update cart item service is not initialized")
	}

	return updateCartItemServiceStruct
}

func (s *updateCartService) ValidateRequest() *errors.ServerError {
	return nil
}

func (s *updateCartService) ProcessRequest() *errors.ServerError {
	return nil
}
