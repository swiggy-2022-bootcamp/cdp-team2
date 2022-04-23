package dao

import "github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"

//Interface for DAO
//go:generate mockgen --destination=../mocks/mock_dao/dao.go github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao IDao
type IDao interface {
	GetByID(string) (*models.Order, error)
	GetAll() ([]models.Order, error)
	Create(models.Order) (*models.Order, error)
	GetByStatus(int) ([]models.Order, error)
	GetByCustomerId(string) ([]models.Order, error)
	UpdateByID(string, models.Order) (*models.Order, error)
	DeleteByID(string) error
}
