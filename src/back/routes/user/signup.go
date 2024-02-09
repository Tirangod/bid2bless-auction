package user

import (
	"bid2bless/src/back/routes/route"

	"github.com/gofiber/fiber/v2"
)

func UserSignupRoute() *route.RouteData {
	return &route.RouteData{
		fiber.MethodPost,
		"/",
		func(c *fiber.Ctx) error {
			return nil
		},
	}
}
