package request

import (
	"user-services/actions/updatepwdbytoken"

	"github.com/gofiber/fiber/v2"
)

func UpdatePwdRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	// Create /login route group.
	group := router.Group("/updatepwd")

	// Register POST /login for customer login.
	group.Post("/", updatepwdbytoken.HandleResetPasswordByToken)
}
