package customerlogin

import (
	"fmt"
	"user-services/gtservices/responsex"
	"user-services/gtservices/utils"

	"github.com/gofiber/fiber/v2"
)

type CustomerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleCustomerLogin(c *fiber.Ctx) error {
	var input CustomerLoginRequest

	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, fmt.Sprintf("Invalid JSON payload: %v", err))
	}

	if input.Email == "" || input.Password == "" {
		return responsex.BadRequest(c, "Email and password are required")
	}
	if !utils.IsValidEmail(input.Email) {
		return responsex.BadRequest(c, "Invalid email format")
	}

	token, err := checkCustomerLogin(input.Email, input.Password)
	if err != nil {
		return responsex.GTError(c, fiber.StatusUnauthorized, fmt.Sprintf("Login failed: %v", err))
	}

	return responsex.GTSuccess(c, "Login successful", fiber.Map{
		"token": token,
	})
}
