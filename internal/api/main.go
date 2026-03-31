package api

import (
	"domains/internal/api/csv"
	"domains/internal/api/domain"
	"domains/internal/api/sheets"
	"log"

	"github.com/gofiber/fiber/v3"
)

func Startup() {
	f := fiber.New()
	var api = f.Group("/api")

	sheets.Handler(api)
	csv.Handler(api)
	domain.Handler(api)

	log.Fatal(f.Listen(":3000"))
}
