package main

import "github.com/gofiber/fiber/v2"

func main() {
	app:= fiber.New()

	// Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Start server
	if err:= app.Listen(":5050"); err != nil {
		panic(err)
	}
}