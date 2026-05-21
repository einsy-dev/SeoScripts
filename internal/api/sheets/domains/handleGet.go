package domains

import (
	"domains/internal/app"
	"domains/internal/models"
	"domains/pkg/csvParser"
	"maps"
	"slices"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm/clause"
)

func handleGet(csv *csvParser.CsvItem) error {
	var csvDoms = slices.Collect(maps.Keys(csv.Rows))
	var dbDoms []models.Domain

	err := app.DB.Preload(clause.Associations).Preload("Outreach.Contact").Where("domain in ?", csvDoms).Find(&dbDoms).Error

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	for _, csvDom := range csvDoms {
		dIndex := slices.IndexFunc(dbDoms, func(el models.Domain) bool {
			return el.Domain == csvDom
		})

		if dIndex == -1 {
			continue
		}

		var dom = dbDoms[dIndex].ToMap()
		if dom != nil {
			csv.UpdateMap(dom)
		}
	}
	return nil
}
