package middleware

import (
	"domains/internal/models"
	"domains/internal/utils"
	"domains/pkg/csvParser"
	"domains/pkg/linkParser"
	"slices"

	"github.com/gofiber/fiber/v3"
)

func Sheets() fiber.Handler {
	return func(c fiber.Ctx) error {
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

			var contact *models.Contact
			contact.ReplaceHeaders(&flatH)

			body.Header = [][]any{flatH}
		}

		csv, err := csvParser.Parse(slices.Concat(body.Header, body.Data), csvParser.Options{})
		csv.FormatRows(func(row string) string {
			return linkParser.Domain(row)
		})

		c.Locals("csv", csv)
		return c.Next()
	}
}
