package handlers

import "github.com/gofiber/fiber/v3"

func App(f *fiber.App) {
	site := f.Group("/app")

	site.Get("/header", func(c fiber.Ctx) error {
		return nil
	})
}
