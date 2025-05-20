package resetpassword

import (
	"fmt"

	"user-services/gtservices/jwtgenx"
	"user-services/gtservices/responsex"
	"user-services/gtservices/utils"

	"github.com/gofiber/fiber/v2"
)

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func HandleResetPassword(c *fiber.Ctx) error {
	var input ResetPasswordRequest

	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, fmt.Sprintf("Invalid request payload: %v", err))
	}

	if input.Email == "" || input.OldPassword == "" || input.NewPassword == "" {
		return responsex.BadRequest(c, "All fields are required")
	}

	if !utils.IsValidEmail(input.Email) {
		return responsex.BadRequest(c, "Invalid email format")
	}
	if !utils.IsValidPassword(input.NewPassword) {
		return responsex.BadRequest(c, "New password must be at least 6 characters, include upper and lower case letters, a number, and a special character (!,@,#,$,%,^)")
	}

	// ✅ Extract Bearer token
	authHeader := c.Get("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		return responsex.GTError(c, fiber.StatusUnauthorized, "Authorization header missing or invalid")
	}
	tokenString := authHeader[7:]

	// ✅ Parse and validate JWT
	claims, err := jwtgenx.ParseToken(tokenString)
	if err != nil {
		return responsex.GTError(c, fiber.StatusUnauthorized, "Invalid or expired token")
	}

	// ✅ Check that token email matches input email
	if claims.Email != input.Email {
		return responsex.BadRequest(c, "Token email does not match the request email")
	}

	// Perform reset
	error := ResetUserPassword(input.Email, input.OldPassword, input.NewPassword)
	if error != nil {
		return responsex.GTError(c, fiber.StatusUnauthorized, fmt.Sprintf("Failed to reset password: %v", error))
	}

	return responsex.GTSuccess(c, "Password updated successfully", nil)
}
