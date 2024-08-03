package api

import (
	"log"

	"github.com/dunky-star/ecommerce-app-golang/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {

	app := fiber.New()

	log.Fatal(app.Listen(config.ServerPort))

}
