package handlers

import (
	"domains/internal/app"
	"domains/internal/models"
	"domains/internal/utils"
	"encoding/json"
	"fmt"
	"strings"

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

	site.Get("/", func(c fiber.Ctx) error {
		var arr [][]any
		err := json.Unmarshal(c.Body(), &arr)
		if err != nil {
			fmt.Println(err)
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}

		var val = utils.Values{}
		val.New(arr)

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
			for k := range val.Cols {
				route := strings.Split(k, ".")

				if len(route) == 1 {
					val.Set(domain["Domain"].(string), k, domain[route[0]])
				} else if domain[route[0]] != nil {
					val.Set(domain["Domain"].(string), k, domain[route[0]].(map[string]any)[route[1]])
				}
			}

		}
		j, err := json.MarshalIndent(val.Data, "", " ")
		return c.Send(j)
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
