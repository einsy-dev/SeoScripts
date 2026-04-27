package account

import (
	"domains/internal/app"
	"domains/internal/models"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	group := f.Group("/account")

	group.Get("/:domain", func(c fiber.Ctx) error {
		// var param = c.Params("domain")
		return c.Status(fiber.StatusOK).JSON(nil)
	})

	group.Post("/:domain", func(c fiber.Ctx) error {
		var param = c.Params("domain")

		d := models.Domain{}
		err := app.DB.Where("domain = ?", param).Preload("Accounts").First(&d).Error
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Domain not found")
		}

		a := models.Account{}
		if err := c.Bind().Body(&a); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
		}

		err = app.DB.Model(&d).Association("Accounts").Append(&a)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(a)
	})
}
