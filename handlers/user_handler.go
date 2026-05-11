package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/database"
	"github.com/kgermando/optimatincorporation-api/models"
	"github.com/kgermando/optimatincorporation-api/utils"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": users})
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	if err := database.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(fiber.Map{"data": user})
}

func CreateUser(c *fiber.Ctx) error {
	type createUserReq struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	var req createUserReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not hash password"})
	}
	user := models.User{Name: req.Name, Email: req.Email, Password: hashed, Role: req.Role}
	if user.Role == "" {
		user.Role = "editor"
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := database.DB.First(&user, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	type updateReq struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}
	var req updateReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	database.DB.Save(&user)
	return c.JSON(fiber.Map{"data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	if err := database.DB.Delete(&models.User{}, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
