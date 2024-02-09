package user

import (
	"github.com/gofiber/fiber/v2"
)

func UserSignupRoute() *RouteData {
	return &RouteData{
		fiber.MethodPost,
		"/",
		func(c *fiber.Ctx) error {
			return nil
		},
	}
}
