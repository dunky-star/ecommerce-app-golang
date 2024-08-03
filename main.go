package main

import (
	"log"

	"github.com/dunky-star/ecommerce-app-golang/config"
	"github.com/dunky-star/ecommerce-app-golang/internal/api"
)

// appConcept "github.com/dunky-star/ecommerce-app-golang/concept"

func main() {

	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)

	// Routes
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, Fiber!!!")
	// })

	//appConcept.MasterConcepts()
	//appConcept.FuncConcepts()

}
