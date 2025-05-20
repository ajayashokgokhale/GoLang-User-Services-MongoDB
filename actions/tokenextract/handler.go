package tokenextract

import (
	"fmt"

	"user-services/gtservices/responsex"

	"github.com/gofiber/fiber/v2"
)

// CustomerLoginRequest defines the structure of the JSON payload for login requests.
type TokenExtractRequest struct {
	Token string
}

func HandleTokenExtract(c *fiber.Ctx) error {
	var input TokenExtractRequest

	// Parse JSON payload from request body.
	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, fmt.Sprintf("Invalid JSON payload: %v", err))
	}

	// Validate required fields and formats.
	if input.Token == "" {
		return responsex.BadRequest(c, "Token is required")
	}

	// Verify credentials and extract generate token.
	email, err := ExtractEmailFromToken(input.Token)
	if err != nil {
		return responsex.GTError(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to extract token: %v", err))
	}
	if email == "" {
		return responsex.GTError(c, fiber.StatusUnauthorized, "Failed to extract token")
	}

	// Return success response with JWT token.
	return responsex.GTSuccess(c, "Toen extract successfully", fiber.Map{
		"email": email,
	})
}
