package newtoken

import (
	"fmt"

	"user-services/gtservices/jwtgenx"

	"github.com/golang-jwt/jwt/v5"
)

// Claims defines the structure for JWT claims, including the user's email.
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func createToken(email string) (string, error) {
	// Validate inputs.
	if email == "" {
		return "", fmt.Errorf("email cannot be empty")
	}

	// Generate JWT token.
	token, err := jwtgenx.GenerateToken(email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
