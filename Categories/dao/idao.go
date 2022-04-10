package dao

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
)

//Interface for DAO
type IDao interface {
	GetByID() (*models.IModel, error)
	GetAll() ([]models.IModel, error)

	Create([]models.IModel) ([]models.IModel, error)
	UpdateByID(int, models.IModel) (*models.IModel, error)
	DeleteByID(...int) []error
}
