package routes

import "github.com/gofiber/fiber/v2"

const defualtIndexPage = "home.html"

func SetupRoutes(app *fiber.App) {
	app.Static("", "src/front", fiber.Static{
		Index: defualtIndexPage,
	})
}
