package services

import "github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"

//interface for category service
//go:generate mockgen --destination=../mocks/mock_services/service.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/services IService
type IService interface {
	GetByID(int) (*models.Order, error)
	//GetAll() ([]models.Order, error)
	Create(models.Order) (*models.Order, error)
	//UpdateByID(int, models.Order) (*models.Order, error)
	//DeleteByID(int) error
	//DeleteMultiple([]int) []error
}
