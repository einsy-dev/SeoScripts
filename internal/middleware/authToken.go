package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthToken() fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or malformed JWT",
			})
		}

		token = strings.TrimPrefix(token, "Bearer ")

		if token != "147852" {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}
}
