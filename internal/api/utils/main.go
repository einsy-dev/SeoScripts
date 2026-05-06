package utils

import (
	"domains/internal/api/utils/csv"
	"domains/internal/api/utils/domain"

	"github.com/gofiber/fiber/v3"
)

func Handler(app fiber.Router) {
	var utils = app.Group("/utils")

	csv.Handler(utils)
	domain.Handler(utils)
}
