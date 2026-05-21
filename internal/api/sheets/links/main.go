package links

import (
	"domains/pkg/csvParser"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	links := f.Group("/links")

	links.Post("/update", func(c fiber.Ctx) error {
		csv := c.Locals("csv").(*csvParser.CsvItem)
		err := handleCreate(csv)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Err handlepost")
		}
		return c.Redirect().Status(fiber.StatusTemporaryRedirect).To("get")
	})
}
