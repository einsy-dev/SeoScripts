package sheets

import (
	"domains/internal/api/sheets/domains"
	"domains/internal/api/sheets/links"
	"domains/internal/api/sheets/table"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	sheets := f.Group("/sheets")

	table.Handlrer(sheets)
	domains.Handler(sheets)
	links.Handler(sheets)
}
