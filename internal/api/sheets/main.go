package sheets

import (
	"domains/internal/api/sheets/accounts"
	"domains/internal/api/sheets/domains"
	"domains/internal/api/sheets/groups"
	"domains/internal/api/sheets/links"
	"domains/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	sheets := f.Group("/sheets")

	sheets.Use(middleware.JsonBody()) // simply adds header to request if body is valid json
	sheets.Use(middleware.Redirect()) // depending on headers redirects to one of following handlers

	domains.Handler(sheets)
	links.Handler(sheets)
	accounts.Handler(sheets)
	groups.Handler(sheets)
}
