package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
)

//interface for category service
//go:generate mockgen --destination=../mocks/mock_services/service.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/services IService
type IService interface {
	GetByID(int) (*models.Category, error)
	GetAll() ([]models.Category, error)
	Create(models.Category) (*models.Category, error)
	UpdateByID(int, models.Category) (*models.Category, error)
	DeleteByID(int) error
	DeleteMultiple([]int) []error
}
