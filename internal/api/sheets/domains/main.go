package domains

import (
	"domains/internal/middleware"
	"domains/internal/utils"
	"domains/pkg/csvParser"
	"slices"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	domains := f.Group("/domains")

	domains.Use(func(c fiber.Ctx) error {
		var body = struct {
			Header [][]any
			Data   [][]any
		}{Header: [][]any{}, Data: [][]any{}}

		err := c.Bind().Body(&body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid body")
		}

		// flat headers
		if len(body.Header) > 1 {
			var flatH = utils.FlatCsv(body.Header)
			body.Header = [][]any{flatH}
		}

		csv, err := csvParser.Parse(slices.Concat(body.Header, body.Data), csvParser.Options{})
		c.Locals("csv", csv)
		return c.Next()
	})

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
			return c.Status(fiber.StatusBadRequest).SendString("Err handleGet")
		}
		return c.Status(fiber.StatusAccepted).JSON(csv.Value)
	})
}
