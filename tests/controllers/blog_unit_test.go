package test_controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-clean-arch/constants"
	"go-clean-arch/interfaces/controllers"
	"go-clean-arch/models"
	"go-clean-arch/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestCreateBlogFailed(t *testing.T) {
	blogData := &models.CreateBlogDTO{}

	blogUsecaseMock := mocks.NewBlogUsecase(t)
	blogUsecaseMock.On("CreateBlog", blogData).Return(constants.StatusInternalServerError, errors.New("Something Failed"))

	blogController := controllers.NewBlogController(blogUsecaseMock)

	app := gin.New()
	app.POST(
		"/blogs",
		func(ctx *gin.Context) {
			ctx.Set("blogData", blogData)
		},
		blogController.CreateNewBlog,
	)

	requestBody, _ := json.Marshal(&models.CreateBlogDTO{
		Title:   "Test",
		Content: "Test",
	})
	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	assert.Equal(t, res.Code, constants.StatusInternalServerError, "Status should be 500")

	blogUsecaseMock.AssertExpectations(t)
}

func TestCreateBlogSuccess(t *testing.T) {
	blogData := &models.CreateBlogDTO{}

	blogUsecaseMock := mocks.NewBlogUsecase(t)
	blogUsecaseMock.On("CreateBlog", blogData).Return(constants.StatusCreated, nil)

	blogController := controllers.NewBlogController(blogUsecaseMock)

	app := gin.New()
	app.POST(
		"/blogs",
		func(ctx *gin.Context) {
			ctx.Set("blogData", blogData)
		},
		blogController.CreateNewBlog,
	)

	requestBody, _ := json.Marshal(&models.CreateBlogDTO{
		Title:   "Test",
		Content: "Test",
	})
	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	app.ServeHTTP(res, req)

	assert.Equal(t, res.Code, constants.StatusCreated, "Status should be 201")

	blogUsecaseMock.AssertExpectations(t)
}
