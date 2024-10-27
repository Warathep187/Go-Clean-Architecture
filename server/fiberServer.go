package server

import (
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/database"
	"go-clean-arch/interfaces/controllers"
	"go-clean-arch/interfaces/middlewares"
	"go-clean-arch/repositories"
	"go-clean-arch/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conf *config.Config
}

func NewServer(conf *config.Config, db database.Database) Server {
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

	s.initialBlogHttpHandler()
	s.initialUserHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Listen(serverUrl)
}

func (s *fiberServer) initialBlogHttpHandler() {
	blogRepo := repositories.NewBlogRepository(s.db)
	userRepo := repositories.NewUserRepository(s.db)
	blogUsecase := usecases.NewBlogUsecase(blogRepo, userRepo)
	blogController := controllers.NewBlogController(blogUsecase)

	blogRouteGroup := s.app.Group("/blogs")
	blogRouteGroup.Post("", middlewares.ValidateBlogData, blogController.CreateNewBlog)
	blogRouteGroup.Get("", blogController.GetAllBlogs)
}

func (s *fiberServer) initialUserHttpHandler() {
	userRepo := repositories.NewUserRepository(s.db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)

	userRouteGroup := s.app.Group("/users")
	userRouteGroup.Post("/register", userController.RegisterUser)
}
