package main

import (
	"bid2bless/src/back/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})
	
	routes.SetupRoutes(app)

	err := app.Listen(":3000")
	log.Fatal(err)
}
