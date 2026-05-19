package links

import (
	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	links := f.Group("/links")

	links.Post("/get", func(c fiber.Ctx) error {
		return nil
	})
	links.Post("/update", func(c fiber.Ctx) error {
		return nil
	})
}
