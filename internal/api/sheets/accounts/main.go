package accounts

import (
	"github.com/gofiber/fiber/v3"
)

func Handler(f fiber.Router) {
	accounts := f.Group("/accounts")

	accounts.Use()

}
