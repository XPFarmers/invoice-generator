package http

import (
	"io/fs"
	"net/http"

	"github.com/XPFarmers/invoice-generator/config"
	"github.com/XPFarmers/invoice-generator/invoices"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewServer(repo *invoices.Repository, config config.Config, staticFiles fs.FS) *fiber.App {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(cors.New())

	// Serve the compile html
	// app.Static("/", "./html")
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(staticFiles),
	}))

	app.Get("/metrics", monitor.New())
	// app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	app.Get("/config", func(c *fiber.Ctx) error {
		return c.JSON(config)
	})

	app.Get("/invoices", func(c *fiber.Ctx) error {
		invoices, err := repo.GetAll()
		if err != nil {
			return c.Status(500).SendString("Could not fetch invoices")
		}
		return c.JSON(invoices)
	})

	app.Get("/invoices/archived", func(c *fiber.Ctx) error {
		archived, err := repo.GetArchived()
		if err != nil {
			return c.Status(500).SendString("Failed to fetch archived")
		}
		return c.JSON(archived)
	})

	app.Post("/invoice", func(c *fiber.Ctx) error {
		body := c.Body()

		invoice, err := invoices.GenerateInvoice(body)
		if err != nil {
			return c.Status(400).SendString("Invalid input: " + err.Error())
		}

		if err := repo.Create(invoice); err != nil {
			return c.Status(500).SendString("Failed to save invoice: " + err.Error())
		}

		return c.Status(201).JSON(invoice)
	})

	// app.Post("/lineItems", func(c *fiber.Ctx) error {
	// var items []invoices.LineItem

	// if err := c.BodyParser(&items); err != nil {
	// 	fmt.Println(err)
	// 	return c.Status(400).SendString("Invalid input: " + err.Error())
	// }

	// // Generate the invoice with deposit, balance, total, etc.
	// invoice := invoices.GenerateInvoice(items)

	// // Optional: Add metadata (you can enhance this later)
	// invoice.ClientName = "Auto Generated"
	// invoice.IssueDate = "2024-06-27"

	// // Save using the repository
	// if err := repo.Create(invoice); err != nil {
	// 	return c.Status(500).SendString("Failed to save invoice: " + err.Error())
	// }

	// return c.Status(201).JSON(invoice)
	// 	fmt.Println(string(c.Body()))

	// 	return c.Status(200).SendString("Success")
	// })

	return app
}
