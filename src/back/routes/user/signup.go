package user

import (
	"bid2bless/src/back/routes/route"

	"github.com/gofiber/fiber/v2"
)

func UserSignupRoute() *route.RouteData {
	return &route.RouteData{
		Method: fiber.MethodPost,
		Path:   "/",
		Handler: func(c *fiber.Ctx) error {
			return nil
		},
	}
}
