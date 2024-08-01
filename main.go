package main

import (
	"log"

	appConcept "github.com/dunky-star/ecommerce-app-golang/concept"
	appConfig "github.com/dunky-star/ecommerce-app-golang/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	appConfig.LoadAppSettings()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!!!")
	})

	appConcept.MasterConcepts()

	log.Fatal(app.Listen(":9000"))

	// Basic Types: int, float64, string, bool
	// Composite Types: array, slice, map, struct
	// Pointer Types: *
}
