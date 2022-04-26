package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	pb "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/protos/shipping"
	grpc_client "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/services/grpc_client_services"
)

type ShippingAddressService struct {
	Dao dao.IDao
}

func NewShippingAddressService() IService {
	return &ShippingAddressService{
		dao.GetShippingDao(),
	}
}

func (cs *ShippingAddressService) GetByCustomerId(customerId string) ([]models.ShippingAddress, error) {
	return cs.Dao.GetByCustomerId(customerId)
}

//
func (cs *ShippingAddressService) Create(shippingAddress models.ShippingAddress) (*models.ShippingAddress, error) {
	return cs.Dao.Create(shippingAddress)
}

func (cs *ShippingAddressService) SetAddressToOrder(orderId string, addressId int) (*pb.OrderAddressUpdateResponse, error) {
	return grpc_client.SetAddress(orderId, addressId)
}
