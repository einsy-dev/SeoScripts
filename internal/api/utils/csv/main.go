package csv

import (
	"domains/internal/utils"
	"domains/pkg/csvParser"
	"domains/pkg/csvParser/services"
	"domains/pkg/linkParser"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Handler(app fiber.Router) {
	var csv = app.Group("/csv")

	// handle single column csv
	csv.Use(func(c fiber.Ctx) error {
		var body interface{}
		var res interface{}
		var err error

		if err := c.Bind().Body(&body); err != nil {
			body = string(c.Body())
		}

		if m, ok := body.(string); ok {
			res, err = services.CsvRead(m)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("err read csv")
			}
		} else if m, ok := body.([]any); ok {
			nr, _ := utils.Assert[[][][]string](res)
			for _, v := range m {
				csv, err := services.CsvRead(v.(string))
				if err != nil {
					return c.Status(fiber.StatusBadRequest).SendString("err read csv")
				}
				nr = append(nr, csv)
			}
			res = nr
		}

		b, err := json.Marshal(res)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("err Marshal body")
		}

		c.Request().SetBody(b)
		c.Request().Header.SetContentType("application/json")

		return c.Next()
	})

	csv.Get("/parse", func(c fiber.Ctx) error {
		var query = struct {
			keys   []string
			format string
			preset string
		}{keys: []string{}}
		query.keys = strings.Split(c.Query("keys"), ",")
		query.format = c.Query("format")
		query.preset = c.Query("preset")

		var res *csvParser.CsvItem

		var data [][]any
		err := c.Bind().Body(&data)

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Invalid body")
		}

		res, err = csvParser.Parse(data, csvParser.Options{Keys: &query.keys})
		if query.format == "domain" {
			res.FormatRows(func(row string) string {
				return linkParser.Domain(row)
			})
		}

		if query.preset == "seo" {
			res.FormatCols(func(row string) string {
				if k, ok := ColsForamtTable[row]; ok {
					return k
				}
				return row
			})
		}

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}

		return c.Status(fiber.StatusOK).JSON(res.Value)
	})

	csv.Get("/join", func(c fiber.Ctx) error {
		var query = struct {
			keys   []string
			format string
			preset string
		}{keys: []string{}}
		query.keys = strings.Split(c.Query("keys"), ",")
		query.format = c.Query("format")
		query.preset = c.Query("preset")

		var res *csvParser.CsvItem
		var data [][][]any
		err := c.Bind().Body(&data)

		if err != nil {
			c.Status(fiber.StatusBadRequest).SendString("Invalid body")
		}

		if len(data) < 1 {
			return c.Status(fiber.StatusBadRequest).SendString("Length is less than 1")
		}

		for _, v := range data {
			if res == nil {
				res, err = csvParser.Parse(v, csvParser.Options{Keys: &query.keys})
				if query.format == "domain" {
					res.FormatRows(func(row string) string {
						return linkParser.Domain(row)
					})
				}
				continue
			}

			c2, err := csvParser.Parse(v, csvParser.Options{Keys: &query.keys})
			if query.format == "domain" {
				c2.FormatRows(func(row string) string {
					return linkParser.Domain(row)
				})
			}

			res.Join(res, c2)

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(res)
			}
		}

		if query.preset == "seo" {
			res.FormatCols(func(row string) string {
				if k, ok := ColsForamtTable[row]; ok {
					return k
				}
				return row
			})
		}

		return c.Status(fiber.StatusOK).JSON(res.Value)
	})
}
