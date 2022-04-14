package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type AddCartItemService interface {
	ValidateRequest() *errors.ServerError
	ProcessRequest() *errors.ServerError
}

var addCartItemServiceStruct AddCartItemService
var addCartItemServiceOnce sync.Once

type addCartItemService struct {
	config *util.RouterConfig
	//dao    mongodao.MongoDAO
}

func InitAddCartItemService(config *util.RouterConfig) AddCartItemService {
	addCartItemServiceOnce.Do(func() {
		addCartItemServiceStruct = &addCartItemService{
			config: config,
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

func (s *addCartItemService) ProcessRequest() *errors.ServerError {
	return nil
}
