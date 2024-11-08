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

	"github.com/gofiber/fiber/v2"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBlogFailedBadRequest(t *testing.T) {
	blogUsecaseMock := mocks.NewBlogUsecase(t)

	blogController := controllers.NewBlogController(blogUsecaseMock)

	app := fiber.New()
	app.Post("/blogs", blogController.CreateNewBlog)

	requestBody := []byte(`invalid_json`)
	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, res.StatusCode, constants.StatusBadRequest, "Status should be 400")
}

func TestCreateBlogFailed(t *testing.T) {
	blogUsecaseMock := mocks.NewBlogUsecase(t)
	blogUsecaseMock.On("CreateBlog", mock.Anything).Return(constants.StatusInternalServerError, errors.New("Something Failed"))

	blogController := controllers.NewBlogController(blogUsecaseMock)

	app := fiber.New()
	app.Post("/blogs", blogController.CreateNewBlog)

	requestBody, _ := json.Marshal(&models.CreateBlogDTO{
		Title:   "Test",
		Content: "Test",
	})
	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, res.StatusCode, constants.StatusInternalServerError, "Status should be 500")

	blogUsecaseMock.AssertExpectations(t)
}

func TestCreateBlogSuccess(t *testing.T) {
	blogUsecaseMock := mocks.NewBlogUsecase(t)
	blogUsecaseMock.On("CreateBlog", mock.Anything).Return(constants.StatusCreated, nil)

	blogController := controllers.NewBlogController(blogUsecaseMock)

	app := fiber.New()
	app.Post("/blogs", blogController.CreateNewBlog)

	requestBody, _ := json.Marshal(&models.CreateBlogDTO{
		Title:   "Test",
		Content: "Test",
	})
	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, res.StatusCode, constants.StatusCreated, "Status should be 201")

	blogUsecaseMock.AssertExpectations(t)
}
