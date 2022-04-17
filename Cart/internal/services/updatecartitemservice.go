package services

import (
	"sync"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"

	log "github.com/sirupsen/logrus"
)

type UpdateCartItemService interface {
	ValidateRequest(product models.Product) *errors.ServerError
	ProcessRequest(customerId string, product models.Product) *errors.ServerError
}

var updateCartItemServiceStruct UpdateCartItemService
var updateCartItemServiceOnce sync.Once

type updateCartService struct {
	config *util.RouterConfig
	dao    dao.DynamoDAO
}

func InitUpdateCartItemService(config *util.RouterConfig, dao dao.DynamoDAO) UpdateCartItemService {
	updateCartItemServiceOnce.Do(func() {
		updateCartItemServiceStruct = &updateCartService{
			config: config,
			dao:    dao,
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

func (s *updateCartService) ValidateRequest(product models.Product) *errors.ServerError {
	if product.ProductId == "" {
		log.Error("product id can not be empty")
		return &errors.ParametersMissingError
	}

	if product.Quantity == 0 {
		log.Error("product quantity can not be zero as cart item")
		return &errors.ParametersMissingError
	}

	return nil
}

func (s *updateCartService) ProcessRequest(customerId string, product models.Product) *errors.ServerError {
	return s.dao.UpdateCartItem(customerId, product.ProductId, product.Quantity)
}
