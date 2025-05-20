package main

import (
	"log"
	"user-services/pkg/dbx"
	"user-services/request"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbx.Mongoinit()

	app := fiber.New()
	request.RegisterRoutes(app)

	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	log.Fatal(app.Listen(port))
}
