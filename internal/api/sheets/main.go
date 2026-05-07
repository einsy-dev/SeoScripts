package sheets

import (
	"domains/internal/api/sheets/accounts"
	"domains/internal/api/sheets/domains"
	"domains/internal/api/sheets/groups"
	"domains/internal/api/sheets/links"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	sheets := f.Group("/sheets")

	domains.Handler(sheets)
	links.Handler(sheets)
	accounts.Handler(sheets)
	groups.Handler(sheets)
}
