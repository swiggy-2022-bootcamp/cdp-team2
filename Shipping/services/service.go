package services

import "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"

//interface for category service
//go:generate mockgen --destination=../mocks/mock_services/service.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/services IService
type IService interface {
	Create(address models.ShippingAddress) (*models.ShippingAddress, error)
	GetByCustomerId(int) ([]models.ShippingAddress, error)
}
