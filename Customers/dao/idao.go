package dao


import "customers/dao/models"

type IDao interface{
	Create(models.Customer) (models.Customer, error)
	Update(string, models.Customer)(models.Customer,error)
	Get( string)(models.Customer,error)
	GetByEmail(string)(models.Customer,error)
	Delete(string)(bool,error)
}