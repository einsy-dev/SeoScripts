package client

import (
	"bytes"

	"github.com/gofiber/fiber/v3"
)

func render(c fiber.Ctx, pageName string, data any) error {
	c.Set("Content-Type", "text/html; charset=utf-8")

	// Determine the target template string layout matching your architecture
	targetTemplate := pageName + "-full"
	if c.Get("HX-Request") == "true" {
		targetTemplate = pageName + "-partial"
	}

	// 1. Render to a buffer memory slice first to prevent partial flushes on template errors
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, targetTemplate, data); err != nil {
		// If template generation breaks, we can cleanly throw a 500 or 400 response code
		return c.Status(fiber.StatusInternalServerError).SendString("Template Error: " + err.Error())
	}

	// 2. Return the buffered byte string payload back down the standard Fiber pipeline
	return c.Send(buf.Bytes())
}
