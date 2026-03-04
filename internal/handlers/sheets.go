package handlers

import (
	"domains/internal/app"
	"domains/internal/models"
	"domains/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Sheets(f *fiber.App) {
	site := f.Group("/sheets")

	site.Get("/", func(c fiber.Ctx) error {
		var arr [][]any
		err := json.Unmarshal(c.Body(), &arr)
		if err != nil {
			fmt.Println(err)
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}
		var val = utils.Values{}
		val.New(arr)
		val.ToMap()

		return c.SendString("Get")
	})

	site.Post("/", func(c fiber.Ctx) error {
		var arr [][]any
		err := json.Unmarshal(c.Body(), &arr)
		if err != nil {
			fmt.Println(err)
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}
		var val = utils.Values{}
		val.New(arr)
		var m = val.ToMap()

		var data []models.Domain
		for key, val := range m {
			var domain models.Domain
			err := app.DB.Preload(clause.Associations).Where(models.Domain{Domain: key}).FirstOrCreate(&domain).Error
			if err != nil {
				fmt.Println(err)
				return nil
			}

			utils.MapToTarget(val, &domain)
			console(val)
			console(domain)

			err = app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&domain).Error
			if err != nil {
				fmt.Println(err)
				return nil
			}
			data = append(data, domain)
		}
		j, err := json.MarshalIndent(data, "", " ")
		return c.Send(j)
	})
}

func console(v any) {
	j, err := json.MarshalIndent(v, "", " ")
	fmt.Println(string(j), err)
}
