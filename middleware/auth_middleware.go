package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/config"
	"github.com/kgermando/optimatincorporation-api/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	cfg := config.Load()
	claims, err := utils.ParseToken(tokenStr, cfg.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("user_role", claims.Role)
	return c.Next()
}
