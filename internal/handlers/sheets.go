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
		var data []map[string]any = utils.ParseValues(arr)

		for _, d := range data {
			var domain models.Domain
			err := app.DB.Preload(clause.Associations).Where(models.Domain{Domain: d["Domain"].(string)}).FirstOrCreate(&domain).Error
			if err != nil {
				fmt.Println(err)
				return nil
			}
			mapToDomain(d, &domain)
			app.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&domain)
		}

		return c.SendString("GET request")
	})
}

func mapToDomain(m map[string]any, target *models.Domain) {
	temp, _ := json.Marshal(m)
	json.Unmarshal(temp, target)
}

func toString(val any) string {
	switch v := val.(type) {
	case string:
		return v
	case float64:
		return fmt.Sprintf("%.0f", v)
	}
	return ""
}
