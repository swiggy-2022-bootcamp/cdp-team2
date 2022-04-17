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

func (cs *OrderService) GetByID(id int) (*models.Order, error) {
	return cs.Dao.GetByID(id)
}

func (cs *OrderService) GetByStatus(status string) ([]models.Order, error) {
	return cs.Dao.GetByStatus(status)
}

func (cs *OrderService) GetAll() ([]models.Order, error) {
	return cs.Dao.GetAll()
}

//
func (cs *OrderService) Create(cat models.Order) (*models.Order, error) {
	return cs.Dao.Create(cat)
}

//
func (cs *OrderService) UpdateByID(id int, cat models.Order) (*models.Order, error) {
	return cs.Dao.UpdateByID(id, cat)
}

func (cs *OrderService) DeleteByID(id int) error {
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
