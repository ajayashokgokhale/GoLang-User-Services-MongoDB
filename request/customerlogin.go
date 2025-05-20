package request

import (
	"user-services/actions/customerlogin"

	"github.com/gofiber/fiber/v2"
)

func CustomerLoginRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	// Create /login route group.
	group := router.Group("/customerlogin")

	// Register POST /login for customer login.
	group.Post("/", customerlogin.HandleCustomerLogin)
}
