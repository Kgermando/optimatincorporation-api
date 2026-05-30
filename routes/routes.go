package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/handlers"
	"github.com/kgermando/optimatincorporation-api/middleware"
)

func Register(api fiber.Router) {
	// Auth
	api.Post("/auth/login", handlers.Login)

	// Public - Portfolio
	api.Get("/portfolio", handlers.GetPortfolioItems)
	api.Get("/portfolio/:slug", handlers.GetPortfolioItem)

	// Public - Services
	api.Get("/services", handlers.GetServices)
	api.Get("/services/:id", handlers.GetService)

	// Public - Blog
	api.Get("/blog", handlers.GetBlogs)
	api.Get("/blog/:slug", handlers.GetBlog)

	// Public - Team
	api.Get("/team", handlers.GetTeamMembers)
	api.Get("/team/:id", handlers.GetTeamMember)

	// Public - Contacts
	api.Post("/contacts", handlers.CreateContact)

	// Upload (no auth required)
	api.Post("/upload", handlers.UploadFile)
	api.Delete("/upload", handlers.DeleteFile)

	// Protected routes
	p := api.Group("", middleware.AuthMiddleware)

	// Portfolio
	p.Post("/portfolio", handlers.CreatePortfolioItem)
	p.Put("/portfolio/:id", handlers.UpdatePortfolioItem)
	p.Delete("/portfolio/:id", handlers.DeletePortfolioItem)

	// Services
	p.Post("/services", handlers.CreateService)
	p.Put("/services/:id", handlers.UpdateService)
	p.Delete("/services/:id", handlers.DeleteService)

	// Blog
	p.Post("/blog", handlers.CreateBlog)
	p.Put("/blog/:id", handlers.UpdateBlog)
	p.Delete("/blog/:id", handlers.DeleteBlog)

	// Team
	p.Post("/team", handlers.CreateTeamMember)
	p.Put("/team/:id", handlers.UpdateTeamMember)
	p.Delete("/team/:id", handlers.DeleteTeamMember)

	// Contacts
	p.Get("/contacts", handlers.GetContacts)
	p.Put("/contacts/:id/read", handlers.MarkContactRead)
	p.Delete("/contacts/:id", handlers.DeleteContact)

	// Users
	p.Get("/users", handlers.GetUsers)
	p.Get("/users/:id", handlers.GetUser)
	p.Post("/users", handlers.CreateUser)
	p.Put("/users/:id", handlers.UpdateUser)
	p.Delete("/users/:id", handlers.DeleteUser)
}
