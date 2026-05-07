package domain

import (
	"domains/internal/utils"
	"domains/pkg/linkParser"

	"github.com/gofiber/fiber/v3"
)

func Handler(app fiber.Router) {
	var domain = app.Group("/domain")

	domain.Post("/:type", func(c fiber.Ctx) error {
		var param = c.Params("type")
		var body []string
		err := c.Bind().Body(&body)

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Invalid body must be [] of strings")
		}

		switch param {
		case "root":
			utils.ForEach(body, func(el string) string {
				return linkParser.RootDomain(el)
			})
		case "domain":
			utils.ForEach(body, func(el string) string {
				return linkParser.Domain(el)
			})
		case "pathname":
			utils.ForEach(body, func(el string) string {
				return linkParser.PathName(el)
			})
		case "params":
			utils.ForEach(body, func(el string) string {
				return linkParser.Params(el)
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(body)
	})

}
