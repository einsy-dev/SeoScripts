package main

import (
	"domains/internal/handlers"
	"domains/internal/models"
	"log"

	"github.com/gofiber/fiber/v3"
)

type Status struct {
	code    uint
	message string
}

func main() {
	models.Startup()

	f := fiber.New()
	handlers.Sheets(f)

	log.Fatal(f.Listen(":3000"))
}
