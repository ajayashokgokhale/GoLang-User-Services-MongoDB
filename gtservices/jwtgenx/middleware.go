package jwtgenx

import (
	"strings"
	"user-services/gtservices/responsex"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return responsex.Unauthorized(c, "missing or invalid token")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ParseToken(tokenStr)
		if err != nil {
			return responsex.Unauthorized(c, "invalid or expired token")
		}

		// You can set user info in context for later use if needed
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
