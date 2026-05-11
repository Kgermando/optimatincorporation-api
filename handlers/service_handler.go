package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/database"
	"github.com/kgermando/optimatincorporation-api/models"
)

func GetServices(c *fiber.Ctx) error {
	var items []models.Service
	if err := database.DB.Where("published = true").Order("`order` asc, created_at asc").Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

func GetService(c *fiber.Ctx) error {
	var item models.Service
	if err := database.DB.First(&item, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(fiber.Map{"data": item})
}

func CreateService(c *fiber.Ctx) error {
	var item models.Service
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"data": item})
}

func UpdateService(c *fiber.Ctx) error {
	var item models.Service
	if err := database.DB.First(&item, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	database.DB.Save(&item)
	return c.JSON(fiber.Map{"data": item})
}

func DeleteService(c *fiber.Ctx) error {
	if err := database.DB.Delete(&models.Service{}, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
