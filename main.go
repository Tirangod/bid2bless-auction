package main

import (
	"bid2bless/src/database"
	"bid2bless/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var mainLogger *log.Logger = log.New(os.Stdout, "Main: ", log.LstdFlags|log.Lmsgprefix)

// Swagger comments...
func main() {
	godotenv.Load() // Load vars from .env file in root folder

	db := database.New()
	if db.Connect() != nil {
		mainLogger.Fatalln("DB connection failed! Exiting...")
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		// Server configuration
	})

	// app.Use(swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: "./docs/openapi.json",
	// 	Path:     "openapi",
	// 	Title:    "Bid2Bless API Docs",
	// })

	routes.SetupRoutes(app)

	err := app.Listen(os.Getenv("PORT"))
	mainLogger.Fatal(err)
}
