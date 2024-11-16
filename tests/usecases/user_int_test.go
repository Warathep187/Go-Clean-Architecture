package test_usecases

import (
	"go-clean-arch/constants"
	"go-clean-arch/database"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	"go-clean-arch/repositories"
	test_setup "go-clean-arch/tests"
	"go-clean-arch/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterUserUsecaseTestSuite struct {
	suite.Suite
	userRepo    repositories.UserRepository
	userUsecase usecases.UserUsecase
}

func TestRegisterUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterUserUsecaseTestSuite))
}

func (suite *RegisterUserUsecaseTestSuite) SetupTest() {
	configs := test_setup.GetTestConfigs()
	db := database.NewPostgresDatabase(configs)
	suite.userRepo = repositories.NewUserRepository(db)
	suite.userUsecase = usecases.NewUserUsecase(suite.userRepo)
}

func (suite *RegisterUserUsecaseTestSuite) TearDownTest() {
	err := suite.userRepo.DeleteUsers()
	assert.Nil(suite.T(), err)
}

func (suite *RegisterUserUsecaseTestSuite) TestRegisterUserFailed_UsernameAlreadyExists() {
	username := "test"
	suite.userRepo.CreateUser(&entities.CreateUserData{
		Username: username,
		Password: "test",
	})

	httpStatus, err := suite.userUsecase.RegisterUser(&models.CreateUserDto{
		Username: username,
		Password: "test",
	})

	assert.Equal(suite.T(), constants.StatusConflict, httpStatus)
	assert.EqualError(suite.T(), err, "Username already exists")
}

func (suite *RegisterUserUsecaseTestSuite) TestRegisterUserSuccess() {
	httpStatus, err := suite.userUsecase.RegisterUser(&models.CreateUserDto{
		Username: "test",
		Password: "test",
	})

	assert.Equal(suite.T(), constants.StatusCreated, httpStatus)
	assert.Nil(suite.T(), err)
}
