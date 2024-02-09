package main

import (
	"bid2bless/src/back/routes"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	db, err := connectToDB()
	if err != nil {
		fmt.Println(err)
	}

	_ = db

	app := fiber.New(fiber.Config{})

	routes.SetupRoutes(app)

	err = app.Listen(":3000")
	log.Fatal(err)
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/data/.db")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to db")
	return db, nil
}
