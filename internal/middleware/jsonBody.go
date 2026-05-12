package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func JsonBody() fiber.Handler {
	return func(c fiber.Ctx) error {

		if json.Valid(c.Body()) {
			c.Request().Header.SetContentType(fiber.MIMEApplicationJSON)
		}

		return c.Next()

	}
}
