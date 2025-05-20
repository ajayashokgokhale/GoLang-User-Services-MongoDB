package createaccount

import (
	"fmt"
	"user-services/gtservices/responsex"

	"github.com/gofiber/fiber/v2"
)

type CreateAccountRequest struct {
	FirstName string `json:"customer_first_name"`
	LastName  string `json:"customer_last_name"`
	Email     string `json:"customer_email"`
	Password  string `json:"customer_password"`
	Phone     string `json:"customer_phone"`
	DOB       string `json:"customer_dob"`
	Gender    string `json:"customer_gender"`
}

func HandleCreateAccount(c *fiber.Ctx) error {
	var input CreateAccountRequest

	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, fmt.Sprintf("Invalid request: %v", err))
	}

	if input.FirstName == "" || input.LastName == "" || input.Email == "" || input.Password == "" {
		return responsex.BadRequest(c, "First name, Last name, Email, and Password are required.")
	}

	customerID, err := CreateCustomer(input)
	if err != nil {
		return responsex.GTError(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to create account: %v", err))
	}

	return responsex.GTSuccess(c, "Account created successfully", fiber.Map{
		"customer_id": customerID,
	})
}
