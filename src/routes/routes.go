package routes

import (
	"bid2bless/src/controllers"

	"github.com/gofiber/fiber/v2"
)

const defualtIndexPage = "index.html"

func SetupRoutes(app *fiber.App) {
	app.Static("", "public", fiber.Static{
		Index: defualtIndexPage,
	})

	// Create user
	app.Post("/user", func(c *fiber.Ctx) error {
		err := controllers.UserSignup(c.Body())
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
		}
		return c.SendStatus(200)
	})
	
	// Login user
	app.Get("/user", func(c *fiber.Ctx) error {
		return nil
	})
}
