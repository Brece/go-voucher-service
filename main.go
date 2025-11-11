package main

import (
	"bytes"

	"github.com/go-pdf/fpdf"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app:= fiber.New()

	// Routes
	app.Get("/health", getHealth)
	app.Post("/voucher", postVoucher)

	// Start server
	if err:= app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func getHealth(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func postVoucher(c *fiber.Ctx) error {
	// Generate PDF voucher
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 12, "Voucher Content Here")			// Placeholder content

	var buf bytes.Buffer

	if err:= pdf.Output(&buf); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate PDF")
	}

	c.Attachment("voucher.pdf")
	c.Type("pdf")

	return c.Send(buf.Bytes())
}