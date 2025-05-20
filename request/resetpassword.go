package request

import (
	"user-services/actions/resetpassword"

	"github.com/gofiber/fiber/v2"
)

func ResetPasswordRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	// Create /reset route group.
	group := router.Group("/resetpassword")

	// Register POST /reset for password reset.
	group.Post("/", resetpassword.HandleResetPassword)
}
