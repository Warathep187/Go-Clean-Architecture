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

func TestRegisterUserFailedBadRequest(t *testing.T) {
	userUsecaseMock := mocks.NewUserUsecase(t)

	userController := controllers.NewUserController(userUsecaseMock)

	app := fiber.New()
	app.Post("/register", userController.RegisterUser)

	requestBody := []byte(`invalid_json`)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, res.StatusCode, constants.StatusBadRequest, "Status should be 400")
}

func TestRegisterUserFailedInternalServerError(t *testing.T) {
	userUsecaseMock := mocks.NewUserUsecase(t)
	userUsecaseMock.On("RegisterUser", mock.Anything).Return(constants.StatusInternalServerError, errors.New("Something Failed"))

	userController := controllers.NewUserController(userUsecaseMock)

	app := fiber.New()
	app.Post("/register", userController.RegisterUser)

	requestBody, _ := json.Marshal(&models.CreateUserDto{
		Username: "test",
		Password: "test",
	})
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, res.StatusCode, constants.StatusInternalServerError, "Status should be 500")

	userUsecaseMock.AssertExpectations(t)
}

func TestRegisterUserSuccess(t *testing.T) {
	userUsecaseMock := mocks.NewUserUsecase(t)
	userUsecaseMock.On("RegisterUser", mock.Anything).Return(constants.StatusOK, nil)

	userController := controllers.NewUserController(userUsecaseMock)

	app := fiber.New()
	app.Post("/register", userController.RegisterUser)

	requestBody, _ := json.Marshal(&models.CreateUserDto{
		Username: "test",
		Password: "test",
	})
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	res, _ := app.Test(req)

	assert.Equal(t, res.StatusCode, constants.StatusOK, "Status should be 200")
}
