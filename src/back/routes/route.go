package routes

import "github.com/gofiber/fiber/v2"

type RouteData struct {
	Method  string
	Path    string
	Handler func(*fiber.Ctx) error
}

func (rd *RouteData) AddRouteToApp(app *fiber.App) {
	app.Add(rd.Method, rd.Path, rd.Handler)
}
