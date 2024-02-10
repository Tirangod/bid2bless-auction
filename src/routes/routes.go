package routes

import (
	"bid2bless/src/controllers"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const defualtIndexPage = "index.html"

func SetupRoutes(app *fiber.App) {
	app.Static("", "public", fiber.Static{
		Index: defualtIndexPage,
	})

	// Create user
	app.Post("/user", func(c *fiber.Ctx) error {
		var u controllers.UserSignupData
		err := json.Unmarshal(c.Body(), &u)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		fmt.Printf("Login: %s, %s, %d;\n", u.Username, u.Email, u.LoginHash)
		return c.SendStatus(200)
	})
	app
}
