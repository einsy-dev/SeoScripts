package handlers

import (
	"domains/internal/models"
	"domains/internal/utils"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func Sheets(f *fiber.App) {
	site := f.Group("/sheets")
	site.Get("/", func(c fiber.Ctx) error {
		var data [][]any
		err := json.Unmarshal(c.Body(), &data)
		if err != nil {
			fmt.Println(err)
			return c.SendString("Error: json.Unmarshal(c.Body(), &data)")
		}

		p := utils.ParseValues(data)
		res := []models.Domain{}
		for _, v := range p {
			d := mapToDomain(v)
			res = append(res, *d)
		}

		// fmt.Println(res)
		// app.DB.Create(res)
		return c.SendString("GET request")
	})
	site.Post("/", func(c fiber.Ctx) error {
		return c.SendString("POST request")
	})
}

func mapToDomain(m map[string]any) *models.Domain {
	d := models.Domain{}

	comment := toString(m["About"].(map[string]any)["Comment"])
	d.Comment = &comment
	fmt.Println(*d.Comment)

	return &d
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
