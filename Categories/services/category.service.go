package services

import (
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
)

type CategoryService struct {
	Dao dao.IDao
}

func NewCategoryService() IService {
	return &CategoryService{
		dao.GetCategoryDao(),
	}
}

func (cs *CategoryService) GetByID(id int) (*models.Category, error) {
	return cs.Dao.GetByID(id)
}

func (cs *CategoryService) GetAll() ([]models.Category, error) {
	return cs.Dao.GetAll()
}

func (cs *CategoryService) Create(cat models.Category) (*models.Category, error) {
	return cs.Dao.Create(cat)
}

func (cs *CategoryService) UpdateByID(id int, cat models.Category) (*models.Category, error) {
	return cs.Dao.UpdateByID(id, cat)
}

func (cs *CategoryService) DeleteByID(id int) error {
	return cs.Dao.DeleteByID(id)
}

func (cs *CategoryService) DeleteMultiple(ids []int) []error {
	var errorList []error
	for _, id := range ids {
		if err := cs.Dao.DeleteByID(id); err != nil {
			errorList = append(errorList, err)
		}
	}
	return errorList
}
