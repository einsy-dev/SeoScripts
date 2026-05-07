package domains

import (
	"domains/internal/app"
	"domains/internal/models"
	"domains/pkg/csvParser"
	"errors"
	"log"
	"maps"
	"slices"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func handleCreate(csv *csvParser.CsvItem) error {
	var csvDoms = slices.Collect(maps.Keys(csv.Rows))
	var dbDoms []models.Domain

	err := app.DB.Preload(clause.Associations).Where("domain in ?", csvDoms).Find(&dbDoms).Error

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	csvMap, err := csv.ToMap()

	for _, csvDom := range csvDoms {
		dIndex := slices.IndexFunc(dbDoms, func(el models.Domain) bool {
			return el.Domain == csvDom
		})

		if dIndex == -1 {
			// create blank if not exist in database
			err := app.DB.Create(&models.Domain{Domain: csvDom}).Error
			if err != nil {
				log.Println(err)
				continue
			}

			var d models.Domain
			err = app.DB.Preload(clause.Associations).Where("domain = ?", csvDom).First(&d).Error
			if err != nil {
				log.Println(err)
				continue
			}

			dbDoms = append(dbDoms, d)
			dIndex = slices.IndexFunc(dbDoms, func(el models.Domain) bool {
				return el.Domain == csvDom
			})
		}

		if dIndex == -1 {
			log.Println("dIndex still -1 after creating and appending new domain")
			continue
		}

		models.MapToDomain(csvMap[csv.Rows[csvDom]-1], &dbDoms[dIndex])

		err := app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dbDoms[dIndex]).Error // optimize for batch upsert

		if err != nil {
			return errors.New("Error saving data " + err.Error())
		}

	}

	return nil
}
