package link

import (
	"domains/internal/app"
	"domains/internal/models"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	link := f.Group("/link")

	link.Post("/:domain", func(c fiber.Ctx) error {
		var param = c.Params("domain")

		var d models.Domain
		err := app.DB.Where("domain = ?", param).Preload("Links").First(&d).Error

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Domain not found")
		}

		var link = models.Link{}

		if err := c.Bind().Body(&link); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
		}

		err = app.DB.Model(&d).Association("Links").Append(&link)

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Domain not found")
		}

		return c.Status(fiber.StatusOK).JSON(link)
	})

	link.Put("/:id", func(c fiber.Ctx) error {
		var param = c.Params("id")

		var link = models.Link{}

		err := app.DB.Where("id = ?", param).First(&link).Error

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Link not found")
		}

		if c.HasBody() {
			if err := c.Bind().Body(&link); err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
			}
		}

		err = app.DB.Save(&link).Error

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(link)
	})
}
