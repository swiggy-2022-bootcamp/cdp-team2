package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Order/internal/literals"
	mock_services "github.com/swiggy-2022-bootcamp/cdp-team2/Order/mocks/mock_services"
)

func getDummyOrder() *models.Order {
	productDesc := []models.ProductDesc{}
	productDesc = append(productDesc, models.ProductDesc{1, 5, 5, 10, 200})
	return &models.Order{"1", "1", 1, 2, 200, 190, productDesc}
}

func MockJsonPost(c *gin.Context /* the test context */, method string, content interface{}) {
	c.Request.Method = method //"POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	res := getDummyOrder()
	mserv.EXPECT().GetByID("1").Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	ctx.Keys = map[string]interface{}{
		literals.OrderIdKey: res.OrderId,
	}

	cont.GetByID(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := models.Order{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, *res, actual)
}

func TestGetByCustomerId(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	res := []models.Order{*getDummyOrder()}
	mserv.EXPECT().GetByCustomerId("1").Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	ctx.Keys = map[string]interface{}{
		literals.CustomerIdKey: res[0].CustomerId,
	}

	cont.GetByCustomer(ctx)

	actual := []models.Order{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, res, actual)
}

func TestGetByStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	res := []models.Order{*getDummyOrder()}
	mserv.EXPECT().GetByStatus(1).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	ctx.Keys = map[string]interface{}{
		literals.StatusKey: int(res[0].Status),
	}

	cont.GetByStatus(ctx)

	actual := []models.Order{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, res, actual)
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	res := []models.Order{*getDummyOrder()}
	mserv.EXPECT().GetAll().Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	cont.GetAll(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := []models.Order{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, res, actual)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	res := getDummyOrder()
	mserv.EXPECT().Create(models.Order{CustomerId: res.CustomerId, Status: res.Status, AddressId: res.AddressId, TotalPrice: res.TotalPrice, PayedPrice: res.PayedPrice, ProductDesc: res.ProductDesc}).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.OrderIdKey:   models.Order{ProductDesc: res.ProductDesc},
		literals.OrderBodyKey: models.Order{CustomerId: res.CustomerId, Status: res.Status, AddressId: res.AddressId, TotalPrice: res.TotalPrice, PayedPrice: res.PayedPrice, ProductDesc: res.ProductDesc},
	}

	cont.Create(ctx)

	actual := models.Order{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.EqualValues(t, res.CustomerId, actual.CustomerId)
	assert.EqualValues(t, res.Status, actual.Status)
	assert.EqualValues(t, res.TotalPrice, actual.TotalPrice)
	assert.EqualValues(t, res.PayedPrice, actual.PayedPrice)
	assert.EqualValues(t, res.AddressId, actual.AddressId)
	assert.EqualValues(t, res.ProductDesc, actual.ProductDesc)

}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	res := getDummyOrder()
	mserv.EXPECT().UpdateByID(res.OrderId, models.Order{ProductDesc: res.ProductDesc}).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.OrderIdKey:   res.OrderId,
		literals.OrderBodyKey: models.Order{ProductDesc: res.ProductDesc},
	}

	cont.Update(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := models.Order{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, *res, actual)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &OrderController{mserv}

	mserv.EXPECT().DeleteByID("1").Times(1).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.OrderIdKey: "1",
	}

	cont.Delete(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestBindId(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Params = []gin.Param{
		{Key: literals.OrderIdKey, Value: "1"},
	}

	BindId(ctx)
	actual := ctx.GetString(literals.OrderIdKey)
	assert.Equal(t, "1", actual)
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
