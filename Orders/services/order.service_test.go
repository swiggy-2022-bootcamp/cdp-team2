package services

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	mock_dao "github.com/swiggy-2022-bootcamp/cdp-team2/Order/mocks/mock_dao"
)

func getDummyOrder() *models.Order {
	productDesc := []models.ProductDesc{}
	productDesc = append(productDesc, models.ProductDesc{1, 5, 5, 10, 200})
	return &models.Order{"1", "2", 2, 2, 200, 190, productDesc}
}

// func getDummyOrders() *[]models.Order {
// 	productDesc1 := []models.ProductDesc{}
// 	productDesc1 = append(productDesc1, models.ProductDesc{1, 5, 5, 10, 200})

// 	productDesc2 := []models.ProductDesc{}
// 	productDesc2 = append(productDesc2, models.ProductDesc{1, 5, 5, 10, 200})

// 	order := []models.Order{}
// 	order = append(order, models.Order{"1", "2", 2, 2, 200, 190, productDesc1})
// 	order = append(order, models.Order{"2", "2", 2, 2, 200, 190, productDesc2})

// 	return &models.Order{"1", "2", 2, 2, 200, 190, productDesc1}
// }

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}

	res := []models.Order{*getDummyOrder()}

	mdao.EXPECT().GetAll().Times(1).Return(res, nil)

	actual, err := catservice.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestGetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}
	res := getDummyOrder()

	mdao.EXPECT().GetByID(res.OrderId).Times(1).Return(res, nil)

	actual, err := catservice.GetByID(res.OrderId)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestGetByCustomerId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}
	res := []models.Order{*getDummyOrder()}

	mdao.EXPECT().GetByCustomerId(res[0].CustomerId).Times(1).Return(res, nil)

	actual, err := catservice.GetByCustomerId(res[0].CustomerId)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestGetByStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}
	res := []models.Order{*getDummyOrder()}

	mdao.EXPECT().GetByStatus(int(res[0].Status)).Times(1).Return(res, nil)

	actual, err := catservice.GetByStatus(int(res[0].Status))
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}
	res := getDummyOrder()
	toCreate := models.Order{ProductDesc: res.ProductDesc}
	mdao.EXPECT().Create(toCreate).Times(1).Return(res, nil)

	actual, err := catservice.Create(toCreate)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "Category should be equal")
}

func TestUpdateByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}
	res := getDummyOrder()
	mdao.EXPECT().UpdateByID(res.OrderId, *res).Times(1).Return(res, nil)

	actual, err := catservice.UpdateByID(res.OrderId, *res)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestDeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	catservice := &OrderService{mdao}
	mdao.EXPECT().DeleteByID("1").Times(1).Return(nil)

	err := catservice.DeleteByID("1")
	assert.Nil(t, err)
}
