package server

import (
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/database"
	"go-clean-arch/interfaces/controllers"
	"go-clean-arch/interfaces/middlewares"
	databaseRepository "go-clean-arch/repositories/database"
	"go-clean-arch/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conf *config.Config
}

func NewFiberServer(conf *config.Config, db database.Database) Server {
	app := fiber.New()

	return &fiberServer{
		app:  app,
		db:   db,
		conf: conf,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(recover.New())

	s.app.Get("/readyz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	blogRepo := databaseRepository.NewBlogRepository(s.db)
	userRepo := databaseRepository.NewUserRepository(s.db)

	blogUsecase := usecases.NewBlogUsecase(blogRepo, userRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	blogController := controllers.NewBlogController(blogUsecase)
	userController := controllers.NewUserController(userUsecase)

	// blog routes
	blogRouteGroup := s.app.Group("/blogs")
	blogRouteGroup.Post("", middlewares.ValidateBlogData, blogController.CreateNewBlog)
	blogRouteGroup.Get("", blogController.GetAllBlogs)

	// user routes
	userRouteGroup := s.app.Group("/users")
	userRouteGroup.Post("/register", userController.RegisterUser)

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Listen(serverUrl)
}
