package request

import (
	"user-services/actions/createaccount"

	"github.com/gofiber/fiber/v2"
)

func CreateAccountRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	// Create /login route group.
	group := router.Group("/createaccount")

	// Register POST /login for customer login.
	group.Post("/", createaccount.HandleCreateAccount)
}
