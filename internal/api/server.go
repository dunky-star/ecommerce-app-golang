package api

import (
	"log"

	"github.com/dunky-star/ecommerce-app-golang/config"
	"github.com/dunky-star/ecommerce-app-golang/internal/api/rest"
	"github.com/dunky-star/ecommerce-app-golang/internal/api/rest/handlers"
	"github.com/dunky-star/ecommerce-app-golang/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {

	app := fiber.New()

	// DB connection
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("database connection error %v\n", err)
	}

	log.Println("database connected")

	// run migration
	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{}, &domain.Category{}, &domain.Product{})
	if err != nil {
		log.Fatalf("error on runing migration %v", err.Error())
	}

	log.Println("migration was successful")

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRouteHandler(rh)

	log.Fatal(app.Listen(config.ServerPort))

}

func setupRouteHandler(rh *rest.RestHandler) {
	// User handler
	handlers.SetupUserRoute(rh)
	// Server health handler
	handlers.SetupServerRoute(rh)
	// Catalog handler

	// Transaction handler
}
