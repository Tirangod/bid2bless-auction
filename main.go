package main

import (
	"bid2bless/src/database"
	"bid2bless/src/routes"
	"bid2bless/src/server"
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var mainLogger *log.Logger = log.New(os.Stdout)

func main() {
	godotenv.Load() // Load vars from .env file in root folder

	db := database.New()
	if db.Connect() != nil {
		mainLogger.Fatal("DB connection failed! Exiting...")
	}
	defer db.Close()

	// No err check because it checked inside server package
	s := server.New()

	routes.SetupRoutes(s.App)

	// No err check because it checked inside server package
	s.Start()
}
