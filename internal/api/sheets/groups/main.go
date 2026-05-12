package groups

import (
	"domains/internal/utils"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	groups := f.Group("/groups")

	groups.Post("/update", func(c fiber.Ctx) error {

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

		return c.Status(fiber.StatusAccepted).SendString("Nice")
	})
}
