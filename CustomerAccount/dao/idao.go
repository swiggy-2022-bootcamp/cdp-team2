package dao


import "customer-account/dao/models"

type IDao interface{
	Create(models.Account) (models.Account, error)
	Update(string, models.Account)(models.Account,error)
	Get( string)(models.Account2,error)
}