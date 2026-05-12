package client

import (
	"embed"
	"html/template"

	"github.com/gofiber/fiber/v3"
)

//go:embed templates/*
var templates embed.FS

func Startup(f *fiber.App) {
	var client = f.Group("/")
	var tmpl = template.Must(template.ParseFS(templates, "templates/pages/*.html", "templates/components/*.html"))

	client.Get("/", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(c.Response().BodyWriter(), "layout", fiber.Map{"Name": "Aliice"})
		return nil
	})

	client.Get("/update", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(c.Response().BodyWriter(), "update", fiber.Map{"Name": "Aliice"})
		return nil
	})
}
