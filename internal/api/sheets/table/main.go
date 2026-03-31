package table

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

func Handlrer(sheets fiber.Router) {
	var table = sheets.Group("/table")

	table.Get("/header", func(c fiber.Ctx) error {
		var h = []string{"", "", "", "", "Ahrefs", "", "", "", "", "", "Semrush", "", "", "", "", "Majestic", "", ""}
		var b = []string{"Domain", "Type", "Rel", "Comment", "DR", "Traffic", "Age", "Geo", "RefDomains", "OutDomains", "AS", "Traffic", "RefDomains", "OutDomains", "LinkFarm", "TF", "CF", "Topic"}
		header := [][]string{h, b}
		j, err := json.MarshalIndent(header, "", " ")
		if err != nil {
			return c.SendStatus(500)
		}
		return c.Send(j)
	})
}
