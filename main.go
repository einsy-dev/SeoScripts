package main

import (
	"domains/internal/api"
	"domains/internal/models"
	"log"
	"os"

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
	api.Startup()
}
