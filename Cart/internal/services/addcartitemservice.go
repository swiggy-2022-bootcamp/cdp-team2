package services

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
)

type AddCartItemService interface {
	ValidateRequest(product models.Product) *errors.ServerError
	ProcessRequest(customerId string, product models.Product) *errors.ServerError
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

func (s *addCartItemService) ValidateRequest(product models.Product) *errors.ServerError {
	if product.ProductId == "" {
		log.Error("product id can not be empty")
		return &errors.ParametersMissingError
	}

	if product.Quantity == 0 {
		log.Error("product quantity can not be zero")
		return &errors.ParametersMissingError
	}

	return nil
}

func (s *addCartItemService) ProcessRequest(customerId string, product models.Product) *errors.ServerError {
	return s.dao.AddCartItem(customerId, product)
}
