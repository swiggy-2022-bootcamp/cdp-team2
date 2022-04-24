package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/internal/literals"
	mock_services "github.com/swiggy-2022-bootcamp/cdp-team2/Shipping/mocks/mock_services"
)

func getDummyShipping() *models.ShippingAddress {
	return &models.ShippingAddress{1, 2, "Taranjeet", "Singh", "Varanasi", "Zamania", "Ghazipur", "IN", 232331}
}

func TestGetByCustomerId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &ShippingAddressController{mserv}

	res := []models.ShippingAddress{*getDummyShipping()}
	mserv.EXPECT().GetByCustomerId(2).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	ctx.Keys = map[string]interface{}{
		literals.CustomerIdKey: res[0].CustomerId,
	}
	// q := ctx.Request.URL.Query()
	// q.Add("")
	//MockJsonPost(ctx, map[string]interface{}{"foo": "bar"})

	cont.GetByCustomer(ctx)
	//assert.EqualValues(t, http.StatusOK, w.Code)

	actual := []models.ShippingAddress{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, res, actual)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &ShippingAddressController{mserv}

	res := getDummyShipping()
	mserv.EXPECT().Create(models.ShippingAddress{res.AddressId, res.CustomerId, res.FirstName, res.LastName, res.City, res.AddressLine1, res.AddressLine2, res.CountryCode, res.PostCode}).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.AddressBodyKey: models.ShippingAddress{res.AddressId, res.CustomerId, res.FirstName, res.LastName, res.City, res.AddressLine1, res.AddressLine2, res.CountryCode, res.PostCode},
	}

	cont.Create(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := models.ShippingAddress{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, *res, actual)

}

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	Health(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
	actual := w.Body.String()
	assert.Equal(t, literals.Health_OK, actual)
}
