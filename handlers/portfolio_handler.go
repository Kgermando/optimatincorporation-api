package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/database"
	"github.com/kgermando/optimatincorporation-api/models"
	"github.com/kgermando/optimatincorporation-api/utils"
)

func GetPortfolioItems(c *fiber.Ctx) error {
	var items []models.Portfolio
	query := database.DB.Where("published = true").Order("year desc, created_at desc")
	if cat := c.Query("category"); cat != "" {
		query = query.Where("category = ?", cat)
	}
	if err := query.Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

func GetPortfolioItem(c *fiber.Ctx) error {
	var item models.Portfolio
	slug := c.Params("slug")
	if err := database.DB.Where("slug = ? OR id = ?", slug, slug).First(&item).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(fiber.Map{"data": item})
}

func CreatePortfolioItem(c *fiber.Ctx) error {
	var item models.Portfolio
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	item.Slug = utils.Slugify(item.Title)
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"data": item})
}

func UpdatePortfolioItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Portfolio
	if err := database.DB.First(&item, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	item.Slug = utils.Slugify(item.Title)
	database.DB.Save(&item)
	return c.JSON(fiber.Map{"data": item})
}

func DeletePortfolioItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Portfolio{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
