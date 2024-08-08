package api

import (
	"log"
	"net/http"

	"github.com/dunky-star/ecommerce-app-golang/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {

	app := fiber.New()

	app.Get("/health", HealthCheck)

	log.Fatal(app.Listen(config.ServerPort))

}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Server heartbeat is alive!",
	})
}
