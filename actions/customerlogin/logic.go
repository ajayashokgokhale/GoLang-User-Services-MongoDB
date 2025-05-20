package customerlogin

import (
	"context"
	"fmt"
	"time"

	"user-services/gtservices/jwtgenx"
	"user-services/gtservices/utils"
	"user-services/pkg/dbx"

	"go.mongodb.org/mongo-driver/bson"
)

type Customer struct {
	Email    string `bson:"customer_email"`
	Password string `bson:"customer_password"`
}

func checkCustomerLogin(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", fmt.Errorf("email and password are required")
	}

	db, err := dbx.GetMongoDB()
	if err != nil {
		return "", fmt.Errorf("database connection failed: %w", err)
	}
	collection := db.Collection("customers")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var customer Customer
	err = collection.FindOne(ctx, bson.M{"customer_email": email}).Decode(&customer)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	ok, err := utils.VerifyPassword(password, customer.Password)
	if err != nil {
		return "", fmt.Errorf("password verification failed: %w", err)
	}
	if !ok {
		return "", fmt.Errorf("invalid email / password")
	}

	token, err := jwtgenx.GenerateToken(email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
