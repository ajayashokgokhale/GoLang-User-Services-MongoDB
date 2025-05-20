package request

import (
	"user-services/actions/newtoken"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewTokenRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	group := router.Group("/createtoken")

	// Register POST /login for customer login.
	group.Post("/", newtoken.HandleNewToken)
}
