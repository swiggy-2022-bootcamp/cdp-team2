package service

import "customer-account/dao/models"

//interface for category service
//go:generate mockgen --destination=../mocks/mock_services/service.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/services IService
type IService interface {
	CreateService(account models.Account) (models.Account, error)
	ReadService(id_string string) (models.Account2,error)
	UpdateService(id_string string,customer models.Account)(models.Account,error)
}
