package api

import (
	"log"

	"github.com/dunky-star/ecommerce-app-golang/config"
	"github.com/dunky-star/ecommerce-app-golang/internal/api/rest"
	"github.com/dunky-star/ecommerce-app-golang/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {

	app := fiber.New()

	rh := &rest.RestHandler{
		App: app,
	}
	setupRouteHandler(rh)

	log.Fatal(app.Listen(config.ServerPort))

}

func setupRouteHandler(rh *rest.RestHandler) {
	handlers.SetupUserRoute(rh)
	handlers.SetupServerRoute(rh)
}
