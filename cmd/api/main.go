package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/iambyt3z/verity-calculator/internal/handlers"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error loading .env file:", err)
	} else {
		fmt.Println(".env file loaded successfully")
	}
}

func main() {
	loadEnv()

	app := fiber.New()

	handlers.Handler(app)

	fmt.Println("Starting Verity Calculator API Service....")

	app.Listen(":3000")
}
