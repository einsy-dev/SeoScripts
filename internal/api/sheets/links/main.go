package links

import (
	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	domains := f.Group("/links")

	domains.Post("/get", func(c fiber.Ctx) error {
		return nil
	})
	domains.Post("/update", func(c fiber.Ctx) error {
		return nil
	})
}
