package handlers

import (
	"net/http"

	"github.com/dunky-star/ecommerce-app-golang/internal/api/rest"
	"github.com/gofiber/fiber/v2"
)

type ServerHandler struct {
	//
}

func SetupServerRoute(rh *rest.RestHandler) {
	app := rh.App
	// To inject any dependencies later
	handler := ServerHandler{}
	// Public end-points
	app.Get("/health", handler.HealthCheck)

}

func (h *ServerHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Server heartbeat is alive!",
	})
}
