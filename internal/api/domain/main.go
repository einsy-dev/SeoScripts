package domain

import (
	"domains/internal/app"
	"domains/internal/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Handler(f fiber.Router) {
	domain := f.Group("/domain")

	domain.Get("/:domain", func(c fiber.Ctx) error {
		var param = c.Params("domain")

		var d models.Domain
		err := app.DB.Preload(clause.Associations).Where("domain = ?", param).First(&d).Error

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Domain not found")
		}

		return c.Status(fiber.StatusOK).JSON(d)
	})

	domain.Post("/:domain", func(c fiber.Ctx) error {
		var param = c.Params("domain")

		var data models.Domain

		err := app.DB.Preload(clause.Associations).Where("domain = ?", param).FirstOrCreate(&data).Error // Updates existing record fields

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "FirstOrCreate err")
		}

		if err := c.Bind().Body(&data); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
		}

		err = app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&data).Error

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(data)
	})

}
