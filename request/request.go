package request

import (
	"user-services/gtservices/jwtgenx"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Validate app parameter.
	if app == nil {
		panic("fiber app cannot be nil")
	}

	// Create /api route group for all API endpoints.
	api := app.Group("/api")

	// Register public routes.

	UpdatePwdRoutes(api) // Update password by token

	RegisterNewTokenRoutes(api)     // New token generation
	RegisterTokenExtractRoutes(api) // Token extraction

	//MongoDBRoutes(api) // MongoDB routes (for testing)
	CreateAccountRoutes(api) // Create account route
	CustomerLoginRoutes(api) // Customer login route

	// Create protected route group with JWT authentication middleware.
	protected := api.Use(jwtgenx.AuthMiddleware())
	ResetPasswordRoutes(protected) // Password reset (authenticated)

}
