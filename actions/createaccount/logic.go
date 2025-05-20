package createaccount

import (
	"context"
	"fmt"
	"time"

	"user-services/gtservices/utils"
	"user-services/pkg/dbx"

	"go.mongodb.org/mongo-driver/bson"
)

type Customer struct {
	CustomerID string    `bson:"customer_id"`
	FirstName  string    `bson:"customer_first_name"`
	LastName   string    `bson:"customer_last_name"`
	Email      string    `bson:"customer_email"`
	Password   string    `bson:"customer_password"`
	Phone      string    `bson:"customer_phone"`
	DOB        string    `bson:"customer_dob"`
	Gender     string    `bson:"customer_gender"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

func CreateCustomer(req CreateAccountRequest) (string, error) {
	db, err := dbx.GetMongoDB()
	if err != nil {
		return "", fmt.Errorf("DB connection error: %w", err)
	}

	collection := db.Collection("customers")

	// Check for existing email
	filter := bson.M{"customer_email": req.Email}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return "", fmt.Errorf("email check failed: %w", err)
	}
	if count > 0 {
		return "", fmt.Errorf("email already exists")
	}

	passwordHash, err := utils.CreateHashPassword(req.Password)
	if err != nil {
		return "", fmt.Errorf("password hashing failed: %w", err)
	}

	customer := Customer{
		CustomerID: fmt.Sprintf("%d", time.Now().UnixNano()), // or use UUID
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Password:   passwordHash,
		Phone:      req.Phone,
		DOB:        req.DOB,
		Gender:     req.Gender,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), customer)
	if err != nil {
		return "", fmt.Errorf("insert failed: %w", err)
	}

	return customer.CustomerID, nil
}
