package domains

import (
	"domains/internal/middleware"
	"domains/pkg/csvParser"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	domains := f.Group("/domains")

	domains.Post("/get", func(c fiber.Ctx) error {
		csv := c.Locals("csv").(*csvParser.CsvItem)
		err := handleGet(csv)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Err handleGet")
		}
		return c.Status(fiber.StatusAccepted).JSON(csv.Value)
	})

	domains.Use(middleware.AuthToken())
	domains.Post("/update", func(c fiber.Ctx) error {
		csv := c.Locals("csv").(*csvParser.CsvItem)

		err := handleCreate(csv)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Err handlepost")
		}

		return c.Redirect().Status(fiber.StatusTemporaryRedirect).To("get")
	})
}
