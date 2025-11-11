package main

import (
	"github.com/gofiber/fiber/v2"
)

// VoucherRequest represents the expected JSON payload for voucher generation
type Client struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Salutation string `json:"salutation"`
}

type Agency struct {
	ID string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Postcode string `json:"postcode"`
	StreetAndNumber string `json:"streetAndNumber"`
	Country string `json:"country"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	CountryCode string `json:"countryCode"`
}

type Activity struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	StartDate string `json:"startDate"`		// YYYY-MM-DD
	EndDate string `json:"endDate"`			// YYYY-MM-DD
}

type VoucherRequest struct {
	BookingID string `json:"bookingId"`
	Activity Activity `json:"activity"`
	Agency Agency `json:"agency"`
	Client Client `json:"client"`
}

// Mock JSON string request object for testing
/*
{
	"bookingId": "BKG123456",
	"activity": {
		"id": "ACT987654",
		"name": "City Tour",
		"code": "CT001",
		"startDate": "2025-02-05",
		"endDate": "2025-02-10"
	},
	"agency": {
		"id": "AGY555555",
		"name": "Travel Experts",
		"city": "New York",
		"postcode": "10001",
		"streetAndNumber": "123 Main St",
		"country": "USA",
		"phone": "+1-555-1234",
		"email": "contact@travelexperts.com",
		"countryCode": "US"		
	},
	"client": {
		"firstName": "John",
		"lastName": "Doe",
		"salutation": "Mr."
		}
}
*/


func registerRoutes(app *fiber.App) {
	// Routes
	app.Get("/health", getHealth)
	app.Post("/voucher", postVoucher)
}

func getHealth(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}


func postVoucher(c *fiber.Ctx) error {
	var req VoucherRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	// Validate required fields (simplified example)
	if req.BookingID == "" || req.Client.FirstName == "" || req.Client.LastName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing required fields")
	}

	b, err := GeneratePDFVoucher(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate PDF")
	}

	c.Attachment("voucher-" + req.BookingID + ".pdf")
	c.Type("pdf")

	return c.Send(b)
}