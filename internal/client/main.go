package client

import (
	"domains/internal/app"
	"domains/internal/models"
	"embed"
	"html/template"

	"github.com/gofiber/fiber/v3"
)

//go:embed templates/*
var templates embed.FS
var tmpl = template.Must(template.ParseFS(templates, "templates/pages/*.html", "templates/components/*.html"))

type Filter struct {
	Search *string `query:"search"`
	Page   *int    `query:"page"`
	Sort   *string `query:"sort"`
}

func Startup(f *fiber.App) {
	var client = f.Group("/")

	client.Get("/", func(c fiber.Ctx) error {
		var f Filter
		c.Bind().Query(f)

		var doms []models.Domain
		app.DB.First(&doms)

		render(c, "main", doms)
		return nil
	})

}
