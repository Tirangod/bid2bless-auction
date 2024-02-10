package main

import (
	"bid2bless/src/database"
	"bid2bless/src/server"
	"log"
	"os"

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

	server, err := server.NewServer(db)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	// app.Use(swagger.Config{
	// 	BasePath: "/",
	// 	FilePath: "./docs/openapi.json",
	// 	Path:     "openapi",
	// 	Title:    "Bid2Bless API Docs",
	// })

}
