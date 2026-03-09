package main

import (
	"domains/internal/handlers"
	"domains/internal/models"
	"log"
	"os"

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
		godotenv.Write(map[string]string{"DB": os.Getenv("DB")}, ".env")
	}

	if os.Getenv("DB") == "" {
		log.Fatal("env not set")
	}

	models.Startup()

	f := fiber.New()
	handlers.Sheets(f)

	log.Fatal(f.Listen(":3000"))
}
