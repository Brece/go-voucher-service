package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app:= fiber.New()

	registerRoutes(app)

	// Start server
	if err:= app.Listen(":8080"); err != nil {
		panic(err)
	}
}


