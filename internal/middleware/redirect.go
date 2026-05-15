package middleware

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Redirect() fiber.Handler {
	return func(c fiber.Ctx) error {
		var body = struct {
			Header [][]any
			Data   [][]any
		}{Header: [][]any{}, Data: [][]any{}}

		err := c.Bind().Body(&body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid body")
		}

		path := c.Path()
		pathSplit := strings.Split(path, "/")
		param := pathSplit[len(pathSplit)-1]

		switch body.Header[0][0] {
		case "F-001":
			c.Path("/api/sheets/domains/" + param)
		case "F-002":
			c.Path("/api/sheets/links/" + param)
		case "F-003":
			c.Path("/api/sheets/accounts/" + param)
		case "F-004":
			c.Path("/api/sheets/domains/" + param)

		}

		body.Header[0][0] = "" // remove id after redirect

		nBody, _ := json.Marshal(body)

		c.Request().SetBody(nBody)

		return c.Next()
	}
}
