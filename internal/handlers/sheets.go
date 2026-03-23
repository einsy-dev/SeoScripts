package handlers

import (
	"domains/internal/app"
	"domains/internal/models"
	"domains/internal/services"
	"domains/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Sheets(f *fiber.App) {
	site := f.Group("/sheets")

	site.Get("/header", func(c fiber.Ctx) error {
		var h = []string{"", "", "", "", "Ahrefs", "", "", "", "", "", "Semrush", "", "", "", "", "Majestic", "", ""}
		var b = []string{"Domain", "Type", "Rel", "Comment", "DR", "Traffic", "Age", "Geo", "RefDomains", "OutDomains", "AS", "Traffic", "RefDomains", "OutDomains", "LinkFarm", "TF", "CF", "Topic"}
		header := [][]string{h, b}
		j, err := json.MarshalIndent(header, "", " ")
		if err != nil {
			return c.SendStatus(500)
		}
		return c.Send(j)
	})

	site.Post("/get-data", func(c fiber.Ctx) error {

		var arr [][]any
		err := json.Unmarshal(c.Body(), &arr)
		if err != nil {
			fmt.Println(err)
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}

		var val = utils.Values{}
		err = val.New(arr)

		if err != nil {
			fmt.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "Can`t process array. Data is invalid")
		}

		var data []models.Domain
		err = app.DB.Preload(clause.Associations).Where("domain IN ?", val.GetKeyRows()).Find(&data).Error
		if err != nil {
			fmt.Println(err)
			return nil
		}

		for _, item := range data {
			domain, err := utils.StructToMap(item)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			val.SetMap(domain)
		}

		j, err := json.MarshalIndent(val.Data, "", " ")
		return c.Send(j)
	})

	site.Use(services.Auth.Token())
	site.Post("/", func(c fiber.Ctx) error {

		var arr [][]any
		err := json.Unmarshal(c.Body(), &arr)
		if err != nil || len(arr) == 0 {
			fmt.Println(err)
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}
		var val = &utils.Values{}
		val.New(arr)
		var domains = val.Rows

		for d := range domains {
			var domain models.Domain
			err := app.DB.Preload(clause.Associations).Where(models.Domain{Domain: d}).FirstOrCreate(&domain).Error
			if err != nil {
				fmt.Println(err)
				return nil
			}

			var dataMap = val.GetMap(d)
			services.MapToDomain(dataMap, &domain)

			err = app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&domain).Error
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}

		j, err := json.MarshalIndent(val.Data, "", " ")
		return c.Send(j)
	})
}
