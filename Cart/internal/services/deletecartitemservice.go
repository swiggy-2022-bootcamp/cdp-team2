package services

import (
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/errors"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DeleteCartItemService interface {
	ProcessRequest(customerId string, productId string) *errors.ServerError
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

func (s *deleteCartItemService) ProcessRequest(customerId string, productId string) *errors.ServerError {
	conn, err := grpc.Dial("0.0.0.0:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Error("unable to connect to grpc service")
		return &errors.InternalError
	}
	defer conn.Close()

	// check if product id exists or not in product MS
	// client := protos.NewRewardServiceClient(conn)
	// req := &protos.RewardRequest{
	// 	CustomerId: "123",
	// }
	// res, err := client.GetReward(context.Background(), req)
	// if err != nil {
	// 	log.WithError(err).Error("product with productId: ", productId, " does not exists in products inventory")
	// 	return &errors.ProductNotExistError
	// }

	return s.dao.DeleteCartItem(customerId, productId)
}
