package services

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	mock_dao "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/mocks/mock_dao"
)

func getDummyShipping() *models.ShippingAddress {
	return &models.ShippingAddress{1, 2, "Taranjeet", "Singh", "Varanasi", "Zamania", "Ghazipur", "IN", 232331}
}

func TestGetByCustomerId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	shippingService := &ShippingAddressService{mdao}
	res := []models.ShippingAddress{*getDummyShipping()}

	mdao.EXPECT().GetByCustomerId(res[0].CustomerId).Times(1).Return(res, nil)

	actual, err := shippingService.GetByCustomerId(res[0].CustomerId)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "they should be equal")
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mdao := mock_dao.NewMockIDao(ctrl)
	shippingService := &ShippingAddressService{mdao}
	res := getDummyShipping()
	toCreate := models.ShippingAddress{CustomerId: res.CustomerId, FirstName: res.FirstName, City: res.City, AddressLine1: res.AddressLine1, AddressLine2: res.AddressLine2, CountryCode: res.CountryCode, PostCode: res.PostCode}
	mdao.EXPECT().Create(toCreate).Times(1).Return(res, nil)

	actual, err := shippingService.Create(toCreate)
	assert.Nil(t, err)
	assert.Equal(t, res, actual, "Shipping Address should be equal")
}
