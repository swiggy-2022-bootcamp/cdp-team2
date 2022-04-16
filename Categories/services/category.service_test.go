package services

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
	mock_dao "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/mocks/mock_dao"
)

func getDummycat() *models.Category {
	return &models.Category{1, models.CategoryDesc{"test", "test", "test", "test", "test"}}
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &CategoryService{mdao}

	res := []models.Category{*getDummycat()}

	mdao.EXPECT().GetAll().Times(1).Return(res, nil)

	actual, err := catservice.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &CategoryService{mdao}
	res := getDummycat()

	mdao.EXPECT().GetByID(res.CategoryId).Times(1).Return(res, nil)

	actual, err := catservice.GetByID(res.CategoryId)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &CategoryService{mdao}
	res := getDummycat()
	toCreate := models.Category{CategoryDesc: res.CategoryDesc}
	mdao.EXPECT().Create(toCreate).Times(1).Return(res, nil)

	actual, err := catservice.Create(toCreate)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "Category should be equal")
}

func TestUpdateByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &CategoryService{mdao}
	res := getDummycat()
	mdao.EXPECT().UpdateByID(res.CategoryId, *res).Times(1).Return(res, nil)

	actual, err := catservice.UpdateByID(res.CategoryId, *res)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestDeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &CategoryService{mdao}
	mdao.EXPECT().DeleteByID(1).Times(1).Return(nil)

	err := catservice.DeleteByID(1)
	assert.Nil(t, err)
}

func TestDeleteMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &CategoryService{mdao}

	gomock.InOrder(
		mdao.EXPECT().DeleteByID(1).Return(nil),
		mdao.EXPECT().DeleteByID(2).Return(fmt.Errorf("some Error")),
		mdao.EXPECT().DeleteByID(3).Return(nil),
	)

	errors := catservice.DeleteMultiple([]int{1, 2, 3})
	assert.Equal(t, 1, len(errors), "there should be 1 error")
	assert.EqualError(t, errors[0], "some Error", "Error should be same")
}
