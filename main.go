package main

import (
	"domains/internal/api"
	"domains/internal/client"
	"domains/internal/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Status struct {
	code    uint
	message string
	demo    Demo
}

type Demo struct {
	test string
}

//go:generate go run ./gen/main.go Status

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

	api.Startup(f)
	client.Startup(f)

	log.Fatal(f.Listen(":3000"))
}
