package main

import (
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	// Routes
	app.Get("/health", getHealth)
	app.Post("/voucher", postVoucher)
}

func getHealth(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}


func postVoucher(c *fiber.Ctx) error {
	b, err := GeneratePDFVoucher()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate PDF")
	}

	c.Attachment("voucher.pdf")
	c.Type("pdf")

	return c.Send(b)
}