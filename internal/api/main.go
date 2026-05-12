package api

import (
	"domains/internal/api/account"
	"domains/internal/api/domain"
	"domains/internal/api/link"
	"domains/internal/api/sheets"
	"domains/internal/api/utils"

	"github.com/gofiber/fiber/v3"
)

func Startup(f *fiber.App) {
	var api = f.Group("/api")

	utils.Handler(api)
	sheets.Handler(api)

	domain.Handler(api)
	link.Handler(api)
	account.Handler(api)
}
