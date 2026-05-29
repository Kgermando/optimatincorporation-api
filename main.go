package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kgermando/optimatincorporation-api/config"
	"github.com/kgermando/optimatincorporation-api/database"
	"github.com/kgermando/optimatincorporation-api/handlers"
	"github.com/kgermando/optimatincorporation-api/middleware"
	"github.com/kgermando/optimatincorporation-api/seeds"
)

func main() {
	cfg := config.Load()
	database.Connect(cfg)
	seeds.Run(database.DB)

	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50 MB
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	api := app.Group("/api/v1")

	// Public
	api.Post("/auth/login", handlers.Login)
	api.Get("/portfolio", handlers.GetPortfolioItems)
	api.Get("/portfolio/:slug", handlers.GetPortfolioItem)
	api.Get("/services", handlers.GetServices)
	api.Get("/services/:id", handlers.GetService)
	api.Get("/blog", handlers.GetBlogs)
	api.Get("/blog/:slug", handlers.GetBlog)
	api.Get("/team", handlers.GetTeamMembers)
	api.Get("/team/:id", handlers.GetTeamMember)
	api.Post("/contacts", handlers.CreateContact)

	// Protected
	p := api.Group("", middleware.AuthMiddleware)
	p.Post("/portfolio", handlers.CreatePortfolioItem)
	p.Put("/portfolio/:id", handlers.UpdatePortfolioItem)
	p.Delete("/portfolio/:id", handlers.DeletePortfolioItem)

	p.Post("/services", handlers.CreateService)
	p.Put("/services/:id", handlers.UpdateService)
	p.Delete("/services/:id", handlers.DeleteService)

	p.Post("/blog", handlers.CreateBlog)
	p.Put("/blog/:id", handlers.UpdateBlog)
	p.Delete("/blog/:id", handlers.DeleteBlog)

	p.Post("/team", handlers.CreateTeamMember)
	p.Put("/team/:id", handlers.UpdateTeamMember)
	p.Delete("/team/:id", handlers.DeleteTeamMember)

	p.Get("/contacts", handlers.GetContacts)
	p.Put("/contacts/:id/read", handlers.MarkContactRead)
	p.Delete("/contacts/:id", handlers.DeleteContact)

	p.Get("/users", handlers.GetUsers)
	p.Get("/users/:id", handlers.GetUser)
	p.Post("/users", handlers.CreateUser)
	p.Put("/users/:id", handlers.UpdateUser)
	p.Delete("/users/:id", handlers.DeleteUser)

	api.Post("/upload", handlers.UploadFile)
	api.Delete("/upload", handlers.DeleteFile)

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
