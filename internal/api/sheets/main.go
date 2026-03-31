package sheets

import (
	"domains/internal/api/sheets/table"
	"domains/internal/app"
	"domains/internal/models"
	"domains/pkg/csv"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Handler(f fiber.Router) {
	sheets := f.Group("/sheets")
	table.Handlrer(sheets)

	sheets.Use(func(c fiber.Ctx) error {
		var data [][]any
		err := json.Unmarshal(c.Body(), &data)
		if err != nil {
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}
		c.Locals("data", data)
		return c.Next()
	})

	sheets.Post("/get", func(c fiber.Ctx) error {
		fmt.Println(c.Locals("data"))
		return c.SendString("RRRR")
	})

	sheets.Post("/update", func(c fiber.Ctx) error {
		var arrMap = csv.Arr2dToMap(c.Locals("data").([][]any))
		var domains = getDomains(arrMap)

		var dbDomains []models.Domain
		app.DB.Where("domain in ?", domains).Preload(clause.Associations).Find(&dbDomains)

		for i, v := range domains {
			dIndex := slices.IndexFunc(dbDomains, func(el models.Domain) bool {
				return el.Domain == v
			})

			if dIndex == -1 {
				fmt.Println("-1")

				err := app.DB.Create(&models.Domain{Domain: v}).Error
				if err != nil {
					fmt.Println(err)
				}
				var d models.Domain
				err = app.DB.Preload(clause.Associations).Where("domain = ?", v).First(&d).Error
				if err != nil {
					fmt.Println(err)
				}

				dbDomains = append(dbDomains, d)
				// shorten
				dIndex = slices.IndexFunc(dbDomains, func(el models.Domain) bool {
					return el.Domain == v
				})

				if dIndex == -1 {
					fmt.Println("dIndex: -1")
					continue
				}
			}

			models.MapToDomain(arrMap[i], &dbDomains[dIndex])
			app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dbDomains)
		}

		return c.Status(fiber.StatusOK).JSON(dbDomains)
	})
}

func getDomains(data []map[string]any) []string {
	var domains = make([]string, len(data))

	for i, v := range data {
		domains[i] = v["Domain"].(string)
	}

	return domains
}
