package server

import (
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/constants"
	"go-clean-arch/interfaces/controllers"
	"go-clean-arch/interfaces/middlewares"
	"go-clean-arch/repositories"
	"go-clean-arch/usecases"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app  *gin.Engine
	conf *config.Config
}

func NewGinServer(conf *config.Config) Server {
	app := gin.Default()

	return &ginServer{
		app:  app,
		conf: conf,
	}
}

func (s *ginServer) Start() {
	s.app.Use(gin.Recovery())

	s.app.GET("/readyz", func(c *gin.Context) {
		c.String(constants.StatusOK, "OK")
	})

	blogRepo := repositories.NewBlogRepository()
	userRepo := repositories.NewUserRepository()

	blogUsecase := usecases.NewBlogUsecase(blogRepo, userRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	// blogs
	blogRouteGroup := s.app.Group("/blogs")
	blogRouteGroup.POST("", middlewares.ValidateBlogData, blogController.CreateNewBlog)
	blogRouteGroup.GET("", blogController.GetAllBlogs)

	// users
	userRouteGroup := s.app.Group("/users")
	userRouteGroup.POST("/register", userController.RegisterUser)

	s.app.Run(fmt.Sprintf(":%d", s.conf.Server.Port))
}
