package domains

import (
	"domains/internal/app"
	"domains/internal/models"
	"domains/pkg/csv"
	"log"
	"slices"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func HanlePost(data [][]any) {
	var arrMap = csv.Arr2dToMap(data)
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
		err := app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dbDomains[dIndex]).Error // optimize for batch upsert
		if err != nil {
			log.Println("Error saving data", v, err.Error())
			return
		}
	}
	log.Println("Data processed: ", len(dbDomains))
}

func getDomains(data []map[string]any) []string {
	var domains = make([]string, len(data))

	for i, v := range data {
		domains[i] = v["Domain"].(string)
	}

	return domains
}
