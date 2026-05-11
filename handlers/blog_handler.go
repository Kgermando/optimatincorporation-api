package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/database"
	"github.com/kgermando/optimatincorporation-api/models"
	"github.com/kgermando/optimatincorporation-api/utils"
)

func GetBlogs(c *fiber.Ctx) error {
	var items []models.Blog
	if err := database.DB.Where("published = true").Order("created_at desc").Find(&items).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

func GetBlog(c *fiber.Ctx) error {
	var item models.Blog
	slug := c.Params("slug")
	if err := database.DB.Where("slug = ? OR id = ?", slug, slug).First(&item).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	return c.JSON(fiber.Map{"data": item})
}

func CreateBlog(c *fiber.Ctx) error {
	var item models.Blog
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	item.Slug = utils.Slugify(item.Title)
	if err := database.DB.Create(&item).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"data": item})
}

func UpdateBlog(c *fiber.Ctx) error {
	var item models.Blog
	if err := database.DB.First(&item, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	item.Slug = utils.Slugify(item.Title)
	database.DB.Save(&item)
	return c.JSON(fiber.Map{"data": item})
}

func DeleteBlog(c *fiber.Ctx) error {
	if err := database.DB.Delete(&models.Blog{}, "id = ?", c.Params("id")).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
