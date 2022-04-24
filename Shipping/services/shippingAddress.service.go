package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
)

type ShippingAddressService struct {
	Dao dao.IDao
}

func NewShippingAddressService() IService {
	return &ShippingAddressService{
		dao.GetShippingDao(),
	}
}

func (cs *ShippingAddressService) GetByCustomerId(customerId int) ([]models.ShippingAddress, error) {
	return cs.Dao.GetByCustomerId(customerId)
}

//
func (cs *ShippingAddressService) Create(cat models.ShippingAddress) (*models.ShippingAddress, error) {
	return cs.Dao.Create(cat)
}
