package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
)

type OrderService struct {
	Dao dao.IDao
}

func NewOrderService() IService {
	return &OrderService{
		dao.GetOrderDao(),
	}
}

func (cs *OrderService) GetByID(id string) (*models.Order, error) {
	return cs.Dao.GetByID(id)
}

func (cs *OrderService) GetByStatus(status int) ([]models.Order, error) {
	return cs.Dao.GetByStatus(status)
}

func (cs *OrderService) GetByCustomerId(customerId string) ([]models.Order, error) {
	return cs.Dao.GetByCustomerId(customerId)
}

func (cs *OrderService) GetAll() ([]models.Order, error) {
	return cs.Dao.GetAll()
}

//
func (cs *OrderService) Create(order models.Order) (*models.Order, error) {
	return cs.Dao.Create(order)
}

//
func (cs *OrderService) UpdateByID(id string, order models.Order) (*models.Order, error) {
	return cs.Dao.UpdateByID(id, order)
}

func (cs *OrderService) DeleteByID(id string) error {
	return cs.Dao.DeleteByID(id)
}

//
//func (cs *OrderService) DeleteMultiple(ids []int) []error {
//	var errorList []error
//	for _, id := range ids {
//		if err := cs.Dao.DeleteByID(id); err != nil {
//			errorList = append(errorList, err)
//		}
//	}
//	return errorList
//}
