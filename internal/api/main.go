package api

import (
	"domains/internal/api/account"
	"domains/internal/api/csv"
	"domains/internal/api/domain"
	"domains/internal/api/link"
	"domains/internal/api/sheets"
	"domains/internal/middleware"
	"log"

	"github.com/gofiber/fiber/v3"
)

func Startup() {
	f := fiber.New()
	var api = f.Group("/api")

	sheets.Handler(api)

	f.Use(middleware.AuthToken())

	csv.Handler(api)
	domain.Handler(api)
	link.Handler(api)
	account.Handler(api)

	log.Fatal(f.Listen(":3000"))
}
