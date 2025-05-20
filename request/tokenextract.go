package request

import (
	"user-services/actions/tokenextract"

	"github.com/gofiber/fiber/v2"
)

func RegisterTokenExtractRoutes(router fiber.Router) {
	// Validate router parameter.
	if router == nil {
		panic("router cannot be nil")
	}

	group := router.Group("/tokenextract")

	// Register POST /login for customer login.
	group.Post("/", tokenextract.HandleTokenExtract)
}
