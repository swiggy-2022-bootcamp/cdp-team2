package dao

import "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"

//Interface for DAO
//go:generate mockgen --destination=../mocks/mock_dao/dao.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao IDao
type IDao interface {
	GetByCustomerId(int) ([]models.ShippingAddress, error)
	Create(models.ShippingAddress) (*models.ShippingAddress, error)
}
