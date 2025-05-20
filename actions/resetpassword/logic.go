package resetpassword

import (
	"context"
	"fmt"
	"time"

	"user-services/gtservices/utils"
	"user-services/pkg/dbx"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Customer struct {
	Email    string `bson:"customer_email"`
	Password string `bson:"customer_password"`
}

func ResetUserPassword(email, oldPassword, newPassword string) error {
	if email == "" || oldPassword == "" || newPassword == "" {
		return fmt.Errorf("email, old password and new password are required")
	}

	db, err := dbx.GetMongoDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	collection := db.Collection("customers")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find customer by email
	var customer Customer
	err = collection.FindOne(ctx, bson.M{"customer_email": email}).Decode(&customer)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Verify old password
	ok, err := utils.VerifyPassword(oldPassword, customer.Password)
	if err != nil || !ok {
		return fmt.Errorf("invalid old password")
	}

	// Hash new password
	newHash, err := utils.CreateHashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	// Update password
	update := bson.M{
		"$set": bson.M{
			"customer_password": newHash,
			"updated_at":        time.Now(),
		},
	}
	result, err := collection.UpdateOne(ctx, bson.M{"customer_email": email}, update, options.Update())
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("password not updated")
	}

	return nil
}
