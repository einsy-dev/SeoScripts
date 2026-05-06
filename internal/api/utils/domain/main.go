package domain

import (
	"domains/pkg/linkParser"

	"github.com/gofiber/fiber/v3"
)

func Handler(app fiber.Router) {
	var domain = app.Group("/domain")

	domain.Post("/format/:type", func(c fiber.Ctx) error {
		var param = c.Params("type")
		var body []string
		err := c.Bind().Body(&body)

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Invalid body must be [] of strings")
		}

		switch param {
		case "root":
			forEach(body, func(el string) string {
				return linkParser.RootDomain(el)
			})
		case "domain":
			forEach(body, func(el string) string {
				return linkParser.Domain(el)
			})
		case "pathname":
			forEach(body, func(el string) string {
				return linkParser.PathName(el)
			})
		case "params":
			forEach(body, func(el string) string {
				return linkParser.Params(el)
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(body)
	})

}

func forEach(s []string, fn func(el string) string) {
	for i, v := range s {
		s[i] = fn(v)
	}
}
