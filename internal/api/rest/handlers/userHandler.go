package handlers

import (
	"net/http"

	"github.com/dunky-star/ecommerce-app-golang/internal/api/rest"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// svc UserService
}

func SetupUserRoute(rh *rest.RestHandler) {
	app := rh.App

	// Create an instance of user service and inject to handler
	handler := UserHandler{}

	// Public end-points
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private end-points
	app.Get("/verify", handler.Login)
	app.Post("/verify", handler.Login)
	app.Post("/profile", handler.Login)
	app.Get("/profile", handler.Login)

	app.Post("/cart", handler.Login)
	app.Get("/cart", handler.Login)
	app.Post("/order", handler.Login)
	app.Get("/order/:id", handler.Login)

	app.Post("/become-seller", handler.Login)

}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Registered",
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Login successfully",
	})
}
