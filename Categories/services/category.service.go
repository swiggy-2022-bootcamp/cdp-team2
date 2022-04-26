package services

import (
	"context"
	"fmt"
	"log"

	productpb "common/protos/products"

	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
	gc "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/grpc/client"
)

type CategoryService struct {
	Dao           dao.IDao
	productClient *gc.ProductClient
}

func NewCategoryService() (IService, error) {
	pc, err := gc.NewProductClient(context.Background())
	if err != nil {
		return nil, err
	}

	return &CategoryService{
		dao.GetCategoryDao(),
		pc,
	}, nil
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

	prodres, err := cs.productClient.GetProductsByCategoryId(cs.productClient.CtxWithTimeOut(), &productpb.CategoryIDRequest{CategoryID: int64(id)})
	if err != nil {
		log.Printf("[Error] fetching product in category %d . error = %s", id, err.Error())
		return err
	}

	log.Printf("[GRPC call] product resp : %+v", prodres)

	if len(prodres.Products) > 0 {
		e := fmt.Errorf("Category %d has products in it, hence cannot be deleted", id)
		log.Println(e)
		return err
	}

	return cs.Dao.DeleteByID(id)
}

func (cs *CategoryService) DeleteMultiple(ids []int) []error {
	var errorList []error
	for _, id := range ids {

		prodres, err := cs.productClient.GetProductsByCategoryId(cs.productClient.CtxWithTimeOut(), &productpb.CategoryIDRequest{CategoryID: int64(id)})
		if err != nil {
			log.Printf("[Error] fetching product in category %d . error = %s", id, err.Error())
			errorList = append(errorList, err)
			continue
		}

		log.Printf("[GRPC call] product resp : %+v", prodres)

		if len(prodres.Products) > 0 {
			e := fmt.Errorf("Category %d has products in it, hence cannot be deleted", id)
			log.Println(e)
			errorList = append(errorList, err)
			continue
		}

		if err := cs.Dao.DeleteByID(id); err != nil {
			errorList = append(errorList, err)
		}
	}
	return errorList
}
