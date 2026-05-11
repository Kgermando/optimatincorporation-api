package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/database"
	"github.com/kgermando/optimatincorporation-api/models"
)

func CreateContact(c *fiber.Ctx) error {
	var item models.Contact
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if item.Name == "" || item.Email == "" || item.Message == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Name, email and message are required"})
	}
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"data": item})
}

func GetContacts(c *fiber.Ctx) error {
	var items []models.Contact
	if err := database.DB.Order("read asc, created_at desc").Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

func MarkContactRead(c *fiber.Ctx) error {
	if err := database.DB.Model(&models.Contact{}).Where("id = ?", c.Params("id")).Update("read", true).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true})
}

func DeleteContact(c *fiber.Ctx) error {
	if err := database.DB.Delete(&models.Contact{}, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
