package newtoken

import (
	"fmt"

	"user-services/gtservices/responsex"
	"user-services/gtservices/utils"

	"github.com/gofiber/fiber/v2"
)

// CustomerLoginRequest defines the structure of the JSON payload for login requests.
type CustomerLoginRequest struct {
	Email string `json:"email"`
}

func HandleNewToken(c *fiber.Ctx) error {
	var input CustomerLoginRequest

	// Parse JSON payload from request body.
	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, fmt.Sprintf("Invalid JSON payload: %v", err))
	}

	// Validate required fields and formats.
	if input.Email == "" {
		return responsex.BadRequest(c, "Email are required")
	}
	if !utils.IsValidEmail(input.Email) {
		return responsex.BadRequest(c, "Invalid email format")
	}

	// Verify credentials and generate token.
	token, err := createToken(input.Email)
	if err != nil {
		return responsex.GTError(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to create token: %v", err))
	}
	if token == "" {
		return responsex.GTError(c, fiber.StatusUnauthorized, "Failed to create token")
	}

	// Return success response with JWT token.
	return responsex.GTSuccess(c, "Toen generated successfully", fiber.Map{
		"token": token,
	})
}
