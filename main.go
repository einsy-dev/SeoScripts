package main

import (
	"domains/internal/handlers"
	"domains/internal/models"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Status struct {
	code    uint
	message string
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	models.Startup()

	f := fiber.New()
	handlers.Sheets(f)

	log.Fatal(f.Listen(":3000"))
}
