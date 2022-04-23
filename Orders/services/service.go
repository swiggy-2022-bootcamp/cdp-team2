package services

import "github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"

//interface for category service
//go:generate mockgen --destination=../mocks/mock_services/service.go github.com/swiggy-2022-bootcamp/cdp-team2/Order/services IService
type IService interface {
	GetByID(string) (*models.Order, error)
	GetAll() ([]models.Order, error)
	Create(models.Order) (*models.Order, error)
	GetByStatus(int) ([]models.Order, error)
	GetByCustomerId(string) ([]models.Order, error)
	UpdateByID(string, models.Order) (*models.Order, error)
	DeleteByID(string) error
	//DeleteMultiple([]int) []error
}
