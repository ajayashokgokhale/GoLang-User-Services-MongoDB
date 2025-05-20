package updatepwdbytoken

import (
	"context"
	"log"
	"time"
	"user-services/gtservices/jwtgenx"
	"user-services/gtservices/responsex"
	"user-services/gtservices/utils"
	"user-services/pkg/dbx"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func HandleResetPasswordByToken(c *fiber.Ctx) error {
	var input ResetPasswordRequest

	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, "invalid request payload")
	}

	if input.Token == "" || input.NewPassword == "" {
		return responsex.BadRequest(c, "token and new_password are required")
	}

	// ✅ Validate token
	claims, err := jwtgenx.ParseToken(input.Token)
	if err != nil || claims.Email == "" {
		return responsex.GTError(c, fiber.StatusUnauthorized, "invalid or expired token")
	}
	email := claims.Email

	// ✅ Validate password format
	if valid := utils.IsValidPassword(input.NewPassword); !valid {
		return responsex.BadRequest(c, "password must be at least 6 characters, contain upper, lower, number, and special character")
	}

	// ✅ Hash password
	hashedPassword, err := utils.CreateHashPassword(input.NewPassword)
	if err != nil {
		return responsex.GTError(c, fiber.StatusInternalServerError, "failed to hash password")
	}

	// ✅ Get MongoDB database
	db, err := dbx.GetMongoDB()
	if err != nil {
		return responsex.GTError(c, fiber.StatusInternalServerError, "failed to connect to database")
	}
	collection := db.Collection("customers")

	// ✅ Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("email", email)
	log.Println("hashedPassword", hashedPassword)
	// ✅ Prepare update
	filter := bson.M{"customer_email": email}
	update := bson.M{
		"$set": bson.M{
			"customer_password": hashedPassword,
			"updated_at":        time.Now(),
		},
	}

	// ✅ Execute update
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return responsex.GTError(c, fiber.StatusInternalServerError, "failed to update password")
	}

	if result.MatchedCount == 0 {
		return responsex.GTError(c, fiber.StatusNotFound, "customer with this email not found")
	}

	return responsex.GTSuccess(c, "password has been reset successfully", nil)
}
