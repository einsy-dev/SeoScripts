package sheets

import (
	"domains/internal/api/sheets/table"
	"domains/internal/app"
	"domains/internal/middleware"
	"domains/internal/models"
	"domains/internal/utils"
	"domains/pkg/csv"
	"encoding/json"
	"log"
	"maps"
	"slices"
	"strings"

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
		var data = c.Locals("data").([][]any)
		arrRows, arrCols := csv.Get2dArrKeys(data)
		var domains = slices.Collect(maps.Keys(arrRows))

		var dbDomains []models.Domain

		err := app.DB.Preload(clause.Associations).Where("domain in ?", domains).Find(&dbDomains).Error

		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		for dName, rowIndex := range arrRows {
			dIndex := slices.IndexFunc(dbDomains, func(el models.Domain) bool {
				return el.Domain == dName
			})
			if dIndex == -1 {
				continue
			}

			var dom = dbDomains[dIndex]
			domain, err := utils.ToMap(dom)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError)
			}

			for c, colIndex := range arrCols {
				addr := strings.Split(c, ".")
				lvl1 := utils.ToCamelCase(addr[0])

				if len(addr) == 1 {
					if v, ok := domain[lvl1]; ok && v != "" {
						data[rowIndex][colIndex] = v
					}
				} else if m, ok := domain[lvl1].(map[string]any); ok {
					lvl2 := utils.ToCamelCase(addr[1])
					if v, ok := m[lvl2]; ok && v != "" {
						data[rowIndex][colIndex] = v
					}
				}
			}
		}

		return c.Status(fiber.StatusOK).JSON(data)
	})

	sheets.Use(middleware.AuthToken())
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
				err := app.DB.Create(&models.Domain{Domain: v}).Error
				if err != nil {
					log.Println(err)
					continue
				}
				var d models.Domain
				err = app.DB.Preload(clause.Associations).Where("domain = ?", v).First(&d).Error
				if err != nil {
					log.Println(err)
					continue
				}

				dbDomains = append(dbDomains, d)
				// shorten
				dIndex = slices.IndexFunc(dbDomains, func(el models.Domain) bool {
					return el.Domain == v
				})
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
