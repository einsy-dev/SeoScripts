package csv

import "github.com/gofiber/fiber/v3"

func Handler(app fiber.Router) {
	var csv = app.Group("/csv")
	csv.Get("/parse", func(c fiber.Ctx) error {
		return nil
	})
}
