package responsex

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GTSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func GTError(c *fiber.Ctx, code int, message string) error {
	// Validate HTTP status code.
	if code < 100 || code > 599 {
		return fmt.Errorf("invalid HTTP status code: %d", code)
	}
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}

// BadRequest sends a 400 Bad Request response with a JSON error message.
func BadRequest(c *fiber.Ctx, message string) error {
	return GTError(c, fiber.StatusBadRequest, message)
}

// Unauthorized sends a 401 Unauthorized response with a JSON error message.
func Unauthorized(c *fiber.Ctx, message string) error {
	return GTError(c, fiber.StatusUnauthorized, message)
}
