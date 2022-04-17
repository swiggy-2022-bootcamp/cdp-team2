package dao

import "github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"

//Interface for DAO
//go:generate mockgen --destination=../mocks/mock_dao/dao.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao IDao
type IDao interface {
	GetByID(int) (*models.Order, error)
	GetAll() ([]models.Order, error)
	Create(models.Order) (*models.Order, error)
	GetByStatus(string) ([]models.Order, error)
	GetByCustomerId(int) ([]models.Order, error)
	UpdateByID(int, models.Order) (*models.Order, error)
	DeleteByID(int) error
}
