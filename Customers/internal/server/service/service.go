package service

import "customers/dao/models"

//interface for category service
//go:generate mockgen --destination=../mocks/mock_services/service.go github.com/swiggy-2022-bootcamp/cdp-team2/Categories/services IService
type IService interface {
	CreateService( models.Customer) (models.Customer, error)
	ReadService( string) (models.Customer,error)
    ReadByEmailService( string)( models.Customer,error)
	UpdateService( string, models.Customer)(models.Customer,error)
    DeleteService( string)( bool,error)
}
