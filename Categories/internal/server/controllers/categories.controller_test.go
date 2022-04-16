package controllers

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
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/dao/models"
	"github.com/swiggy-2022-bootcamp/cdp-team2/Categories/internal/literals"
	mock_services "github.com/swiggy-2022-bootcamp/cdp-team2/Categories/mocks/mock_services"
)

func getDummycat() *models.Category {
	return &models.Category{1, models.CategoryDesc{"test", "test", "test", "test", "test"}}
}

func MockJsonPost(c *gin.Context /* the test context */, method string, content interface{}) {
	c.Request.Method = method //"POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &CategoryController{mserv}

	res := getDummycat()
	mserv.EXPECT().GetByID(1).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	ctx.Keys = map[string]interface{}{
		literals.CatIdKey: res.CategoryId,
	}
	// q := ctx.Request.URL.Query()
	// q.Add("")
	//MockJsonPost(ctx, map[string]interface{}{"foo": "bar"})

	cont.GetByID(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := models.Category{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, *res, actual)
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &CategoryController{mserv}

	res := []models.Category{*getDummycat()}
	mserv.EXPECT().GetAll().Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	cont.GetAll(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := []models.Category{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, res, actual)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &CategoryController{mserv}

	res := getDummycat()
	mserv.EXPECT().Create(models.Category{CategoryDesc: res.CategoryDesc}).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.CatBodyKey: models.Category{CategoryDesc: res.CategoryDesc},
	}

	cont.Create(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := models.Category{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, *res, actual)
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &CategoryController{mserv}

	res := getDummycat()
	mserv.EXPECT().UpdateByID(res.CategoryId, models.Category{CategoryDesc: res.CategoryDesc}).Times(1).Return(res, nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.CatIdKey:   res.CategoryId,
		literals.CatBodyKey: models.Category{CategoryDesc: res.CategoryDesc},
	}

	cont.Update(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)

	actual := models.Category{}
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Errorf("error unmarshing response json")
	}

	assert.Equal(t, *res, actual)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &CategoryController{mserv}

	mserv.EXPECT().DeleteByID(1).Times(1).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Keys = map[string]interface{}{
		literals.CatIdKey: 1,
	}

	cont.Delete(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestDeleteMultiple(t *testing.T) {
	ctrl := gomock.NewController(t)
	mserv := mock_services.NewMockIService(ctrl)
	cont := &CategoryController{mserv}

	ids := []int{1, 2, 3}
	mserv.EXPECT().DeleteMultiple(ids).Times(1).Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(ctx, "DELETE", ids)

	cont.DeleteMultiple(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestBindCategory(t *testing.T) {
	cat := getDummycat()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(ctx, "POST", *cat)

	BindCategory(ctx)
	actual, exist := ctx.Get(literals.CatBodyKey)
	assert.True(t, exist)
	assert.Equal(t, *cat, actual.(models.Category))
}

func TestBindId(t *testing.T) {

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Params = []gin.Param{
		{Key: literals.CatIdKey, Value: "1"},
	}

	BindId(ctx)
	actual := ctx.GetInt(literals.CatIdKey)
	assert.Equal(t, 1, actual)
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
