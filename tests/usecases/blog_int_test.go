package test_usecases

import (
	"fmt"
	"go-clean-arch/constants"
	"go-clean-arch/database"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	databaseRepository "go-clean-arch/repositories/database"
	test_setup "go-clean-arch/tests"
	"go-clean-arch/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// CREATE BLOG
type CreateBlogUsecaseTestSuite struct {
	suite.Suite
	blogRepo    databaseRepository.BlogRepository
	userRepo    databaseRepository.UserRepository
	blogUsecase usecases.BlogUsecase
}

func TestCreateBlogTestSuite(t *testing.T) {
	suite.Run(t, new(CreateBlogUsecaseTestSuite))
}

func (suite *CreateBlogUsecaseTestSuite) SetupTest() {
	configs := test_setup.GetTestConfigs()
	db := database.NewPostgresDatabase(configs)
	suite.blogRepo = databaseRepository.NewBlogRepository(db)
	suite.userRepo = databaseRepository.NewUserRepository(db)
	suite.blogUsecase = usecases.NewBlogUsecase(suite.blogRepo, suite.userRepo)
}

func (suite *CreateBlogUsecaseTestSuite) TestCreateBlogFailed_UserNotFound() {
	blogData := models.CreateBlogDTO{
		UserID:  1,
		Title:   "Test Blog",
		Content: "Test Content",
	}

	statusCode, err := suite.blogUsecase.CreateBlog(&blogData)
	assert.Equal(suite.T(), constants.StatusNotFound, statusCode)
	assert.Equal(suite.T(), "User not found. Cannot create blog.", err.Error())
}

func (suite *CreateBlogUsecaseTestSuite) TestCreateBlogSuccess() {
	userID := uint(1)
	err := suite.userRepo.CreateUserWithID(userID, &entities.CreateUserData{
		Username: "testuser",
		Password: "testpassword",
	})
	fmt.Println("err", err)
	blogData := models.CreateBlogDTO{
		UserID:  userID,
		Title:   "Test Blog",
		Content: "Test Content",
	}
	httpStatus, err := suite.blogUsecase.CreateBlog(&blogData)
	assert.Equal(suite.T(), constants.StatusCreated, httpStatus)
	assert.Nil(suite.T(), err)

	err = suite.blogRepo.DeleteBlogs()
	assert.Nil(suite.T(), err)
	err = suite.userRepo.DeleteUsers()
	assert.Nil(suite.T(), err)
}

// GET ALL BLOGS
type GetAllBlogsUsecaseTestSuite struct {
	suite.Suite
	blogRepo    databaseRepository.BlogRepository
	blogUsecase usecases.BlogUsecase
}

func TestGetAllBlogsTestSuite(t *testing.T) {
	suite.Run(t, new(GetAllBlogsUsecaseTestSuite))
}

func (suite *GetAllBlogsUsecaseTestSuite) SetupTest() {
	configs := test_setup.GetTestConfigs()
	db := database.NewPostgresDatabase(configs)
	suite.blogRepo = databaseRepository.NewBlogRepository(db)
	suite.blogUsecase = usecases.NewBlogUsecase(suite.blogRepo, nil)
}

func (suite *GetAllBlogsUsecaseTestSuite) TestGetAllBlogsSuccess() {
	err := suite.blogRepo.CreateBlog(&entities.CreateBlogData{
		Title:   "Test Blog",
		Content: "Test Content",
	})
	assert.Nil(suite.T(), err)

	blogs, httpStatus, err := suite.blogUsecase.GetAllBlogs()
	assert.NotNil(suite.T(), blogs)
	assert.Equal(suite.T(), 1, len(blogs))
	assert.Equal(suite.T(), "Test Blog", blogs[0].Title)
	assert.Equal(suite.T(), "Test Content", blogs[0].Content)
	assert.Equal(suite.T(), constants.StatusOK, httpStatus)
	assert.Nil(suite.T(), err)

	err = suite.blogRepo.DeleteBlogs()
	assert.Nil(suite.T(), err)
}
