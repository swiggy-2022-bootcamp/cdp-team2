package dao

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
)

//Interface for DAO
//go:generate mockgen --destination=../mocks/mock_dao/dao.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao IDao
type IDao interface {
	GetByID(int) (*models.Category, error)
	GetAll() ([]models.Category, error)
	Create(models.Category) (*models.Category, error)
	UpdateByID(int, models.Category) (*models.Category, error)
	DeleteByID(int) error
}
